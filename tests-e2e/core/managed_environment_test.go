package core

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	appv1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	managedgitopsv1alpha1 "github.com/redhat-appstudio/managed-gitops/backend-shared/apis/managed-gitops/v1alpha1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redhat-appstudio/managed-gitops/backend-shared/config/db"
	dbutil "github.com/redhat-appstudio/managed-gitops/backend-shared/config/db/util"
	argocdutil "github.com/redhat-appstudio/managed-gitops/backend-shared/util/argocd"
	"github.com/redhat-appstudio/managed-gitops/tests-e2e/fixture"
	appFixture "github.com/redhat-appstudio/managed-gitops/tests-e2e/fixture/application"
	gitopsDeplFixture "github.com/redhat-appstudio/managed-gitops/tests-e2e/fixture/gitopsdeployment"
	"github.com/redhat-appstudio/managed-gitops/tests-e2e/fixture/k8s"
	apps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("GitOpsDeployment Managed Environment E2E tests", func() {

	Context("Create a new GitOpsDeployment targeting a ManagedEnvironment", func() {

		It("should be healthy and have synced status, and resources should be deployed, when deployed with a ManagedEnv", func() {

			Expect(fixture.EnsureCleanSlate()).To(Succeed())

			By("creating the GitOpsDeploymentManagedEnvironment")

			kubeConfigContents, apiServerURL, err := extractKubeConfigValues()
			Expect(err).To(BeNil())

			managedEnv, secret := buildManagedEnvironment(apiServerURL, kubeConfigContents)

			err = k8s.Create(&secret)
			Expect(err).To(BeNil())

			err = k8s.Create(&managedEnv)
			Expect(err).To(BeNil())

			gitOpsDeploymentResource := buildGitOpsDeploymentResource("my-gitops-depl",
				"https://github.com/redhat-appstudio/gitops-repository-template", "environments/overlays/dev",
				managedgitopsv1alpha1.GitOpsDeploymentSpecType_Automated)
			gitOpsDeploymentResource.Spec.Destination.Environment = managedEnv.Name
			gitOpsDeploymentResource.Spec.Destination.Namespace = fixture.GitOpsServiceE2ENamespace
			err = k8s.Create(&gitOpsDeploymentResource)
			Expect(err).To(BeNil())

			By("ensuring GitOpsDeployment should have expected health and status")

			Eventually(gitOpsDeploymentResource, "2m", "1s").Should(
				SatisfyAll(
					gitopsDeplFixture.HaveSyncStatusCode(managedgitopsv1alpha1.SyncStatusCodeSynced),
					gitopsDeplFixture.HaveHealthStatusCode(managedgitopsv1alpha1.HeathStatusCodeHealthy)))

			secretList := corev1.SecretList{}

			k8sClient, err := fixture.GetKubeClient()
			Expect(err).To(BeNil())
			err = k8sClient.List(context.Background(), &secretList, &client.ListOptions{Namespace: dbutil.DefaultGitOpsEngineSingleInstanceNamespace})
			Expect(err).To(BeNil())

			dbQueries, err := db.NewSharedProductionPostgresDBQueries(false)
			Expect(err).To(BeNil())
			defer dbQueries.CloseDatabase()

			mapping := &db.APICRToDatabaseMapping{
				APIResourceType: db.APICRToDatabaseMapping_ResourceType_GitOpsDeploymentManagedEnvironment,
				APIResourceUID:  string(managedEnv.UID),
				DBRelationType:  db.APICRToDatabaseMapping_DBRelationType_ManagedEnvironment,
			}
			err = dbQueries.GetDatabaseMappingForAPICR(context.Background(), mapping)
			Expect(err).To(BeNil())

			argoCDClusterSecret := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      argocdutil.GenerateArgoCDClusterSecretName(db.ManagedEnvironment{Managedenvironment_id: mapping.DBRelationKey}),
					Namespace: dbutil.DefaultGitOpsEngineSingleInstanceNamespace,
				},
			}

			Expect(argoCDClusterSecret).To(k8s.ExistByName())

			By("ensuring the resources of the GitOps repo are successfully deployed")

			componentADepl := &apps.Deployment{
				ObjectMeta: metav1.ObjectMeta{Name: "component-a", Namespace: fixture.GitOpsServiceE2ENamespace},
			}
			componentBDepl := &apps.Deployment{
				ObjectMeta: metav1.ObjectMeta{Name: "component-b", Namespace: fixture.GitOpsServiceE2ENamespace},
			}
			Eventually(componentADepl, "60s", "1s").Should(k8s.ExistByName())
			Eventually(componentBDepl, "60s", "1s").Should(k8s.ExistByName())

			By("deleting the secret and managed environment")
			err = k8s.Delete(&secret)
			Expect(err).To(BeNil())

			err = k8s.Delete(&managedEnv)
			Expect(err).To(BeNil())

			Eventually(argoCDClusterSecret, "60s", "1s").ShouldNot(k8s.ExistByName(),
				"once the ManagedEnvironment is deleted, the Argo CD cluster secret should be deleted as well.")

			app := appv1alpha1.Application{
				ObjectMeta: metav1.ObjectMeta{
					Name:      argocdutil.GenerateArgoCDApplicationName(string(gitOpsDeploymentResource.UID)),
					Namespace: dbutil.GetGitOpsEngineSingleInstanceNamespace(),
				},
			}
			Eventually(app, "60s", "1s").Should(appFixture.HasDestinationField(appv1alpha1.ApplicationDestination{
				Namespace: gitOpsDeploymentResource.Spec.Destination.Namespace,
				Name:      "",
			}), "the Argo CD Application resource's spec.destination field should have an empty environment field")

			By("deleting the GitOpsDeployment")

			err = k8s.Delete(&gitOpsDeploymentResource)
			Expect(err).To(Succeed())

		})
	})
})

// extractKubeConfigValues returns contents of k8s config from $KUBE_CONFIG, plus server api url (and error)
func extractKubeConfigValues() (string, string, error) {

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()

	config, err := loadingRules.Load()
	if err != nil {
		return "", "", err
	}

	context, ok := config.Contexts[config.CurrentContext]
	if !ok || context == nil {
		return "", "", fmt.Errorf("no context")
	}

	cluster, ok := config.Clusters[context.Cluster]
	if !ok || cluster == nil {
		return "", "", fmt.Errorf("no cluster")
	}

	var kubeConfigDefault string

	paths := loadingRules.Precedence
	{

		for _, path := range paths {

			GinkgoWriter.Println("Attempting to read kube config from", path)

			// homeDir, err := os.UserHomeDir()
			// if err != nil {
			// 	return "", "", err
			// }

			_, err = os.Stat(path)
			if err != nil {
				GinkgoWriter.Println("Unable to resolve path", path, err)
			} else {
				// Success
				kubeConfigDefault = path
				break
			}

		}

		if kubeConfigDefault == "" {
			return "", "", fmt.Errorf("unable to retrieve kube config path")
		}
	}

	kubeConfigContents, err := ioutil.ReadFile(kubeConfigDefault)
	if err != nil {
		return "", "", err
	}

	return string(kubeConfigContents), cluster.Server, nil
}

func buildManagedEnvironment(apiServerURL string, kubeConfigContents string) (managedgitopsv1alpha1.GitOpsDeploymentManagedEnvironment, corev1.Secret) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-managed-env-secret",
			Namespace: fixture.GitOpsServiceE2ENamespace,
		},
		Type:       "managed-gitops.redhat.com/managed-environment",
		StringData: map[string]string{"kubeconfig": kubeConfigContents},
	}

	managedEnv := &managedgitopsv1alpha1.GitOpsDeploymentManagedEnvironment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-managed-env",
			Namespace: fixture.GitOpsServiceE2ENamespace,
		},
		Spec: managedgitopsv1alpha1.GitOpsDeploymentManagedEnvironmentSpec{
			APIURL:                   apiServerURL,
			ClusterCredentialsSecret: secret.Name,
		},
	}

	return *managedEnv, *secret
}
