//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDestination) DeepCopyInto(out *ApplicationDestination) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDestination.
func (in *ApplicationDestination) DeepCopy() *ApplicationDestination {
	if in == nil {
		return nil
	}
	out := new(ApplicationDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSource) DeepCopyInto(out *ApplicationSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSource.
func (in *ApplicationSource) DeepCopy() *ApplicationSource {
	if in == nil {
		return nil
	}
	out := new(ApplicationSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeployment) DeepCopyInto(out *GitOpsDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeployment.
func (in *GitOpsDeployment) DeepCopy() *GitOpsDeployment {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitOpsDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentCondition) DeepCopyInto(out *GitOpsDeploymentCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentCondition.
func (in *GitOpsDeploymentCondition) DeepCopy() *GitOpsDeploymentCondition {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentList) DeepCopyInto(out *GitOpsDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitOpsDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentList.
func (in *GitOpsDeploymentList) DeepCopy() *GitOpsDeploymentList {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitOpsDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentManagedEnvironment) DeepCopyInto(out *GitOpsDeploymentManagedEnvironment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentManagedEnvironment.
func (in *GitOpsDeploymentManagedEnvironment) DeepCopy() *GitOpsDeploymentManagedEnvironment {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentManagedEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitOpsDeploymentManagedEnvironment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentManagedEnvironmentList) DeepCopyInto(out *GitOpsDeploymentManagedEnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitOpsDeploymentManagedEnvironment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentManagedEnvironmentList.
func (in *GitOpsDeploymentManagedEnvironmentList) DeepCopy() *GitOpsDeploymentManagedEnvironmentList {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentManagedEnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitOpsDeploymentManagedEnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentManagedEnvironmentSpec) DeepCopyInto(out *GitOpsDeploymentManagedEnvironmentSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentManagedEnvironmentSpec.
func (in *GitOpsDeploymentManagedEnvironmentSpec) DeepCopy() *GitOpsDeploymentManagedEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentManagedEnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentManagedEnvironmentStatus) DeepCopyInto(out *GitOpsDeploymentManagedEnvironmentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentManagedEnvironmentStatus.
func (in *GitOpsDeploymentManagedEnvironmentStatus) DeepCopy() *GitOpsDeploymentManagedEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentManagedEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentRepositoryCredential) DeepCopyInto(out *GitOpsDeploymentRepositoryCredential) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentRepositoryCredential.
func (in *GitOpsDeploymentRepositoryCredential) DeepCopy() *GitOpsDeploymentRepositoryCredential {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentRepositoryCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitOpsDeploymentRepositoryCredential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentRepositoryCredentialList) DeepCopyInto(out *GitOpsDeploymentRepositoryCredentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitOpsDeploymentRepositoryCredential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentRepositoryCredentialList.
func (in *GitOpsDeploymentRepositoryCredentialList) DeepCopy() *GitOpsDeploymentRepositoryCredentialList {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentRepositoryCredentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitOpsDeploymentRepositoryCredentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentRepositoryCredentialSpec) DeepCopyInto(out *GitOpsDeploymentRepositoryCredentialSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentRepositoryCredentialSpec.
func (in *GitOpsDeploymentRepositoryCredentialSpec) DeepCopy() *GitOpsDeploymentRepositoryCredentialSpec {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentRepositoryCredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentRepositoryCredentialStatus) DeepCopyInto(out *GitOpsDeploymentRepositoryCredentialStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentRepositoryCredentialStatus.
func (in *GitOpsDeploymentRepositoryCredentialStatus) DeepCopy() *GitOpsDeploymentRepositoryCredentialStatus {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentRepositoryCredentialStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentSpec) DeepCopyInto(out *GitOpsDeploymentSpec) {
	*out = *in
	out.Source = in.Source
	out.Destination = in.Destination
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentSpec.
func (in *GitOpsDeploymentSpec) DeepCopy() *GitOpsDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentStatus) DeepCopyInto(out *GitOpsDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]GitOpsDeploymentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Sync = in.Sync
	out.Health = in.Health
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentStatus.
func (in *GitOpsDeploymentStatus) DeepCopy() *GitOpsDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentSyncRun) DeepCopyInto(out *GitOpsDeploymentSyncRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentSyncRun.
func (in *GitOpsDeploymentSyncRun) DeepCopy() *GitOpsDeploymentSyncRun {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentSyncRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitOpsDeploymentSyncRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentSyncRunCondition) DeepCopyInto(out *GitOpsDeploymentSyncRunCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentSyncRunCondition.
func (in *GitOpsDeploymentSyncRunCondition) DeepCopy() *GitOpsDeploymentSyncRunCondition {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentSyncRunCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentSyncRunList) DeepCopyInto(out *GitOpsDeploymentSyncRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitOpsDeploymentSyncRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentSyncRunList.
func (in *GitOpsDeploymentSyncRunList) DeepCopy() *GitOpsDeploymentSyncRunList {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentSyncRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitOpsDeploymentSyncRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentSyncRunSpec) DeepCopyInto(out *GitOpsDeploymentSyncRunSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentSyncRunSpec.
func (in *GitOpsDeploymentSyncRunSpec) DeepCopy() *GitOpsDeploymentSyncRunSpec {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentSyncRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsDeploymentSyncRunStatus) DeepCopyInto(out *GitOpsDeploymentSyncRunStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]GitOpsDeploymentSyncRunCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsDeploymentSyncRunStatus.
func (in *GitOpsDeploymentSyncRunStatus) DeepCopy() *GitOpsDeploymentSyncRunStatus {
	if in == nil {
		return nil
	}
	out := new(GitOpsDeploymentSyncRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthStatus) DeepCopyInto(out *HealthStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthStatus.
func (in *HealthStatus) DeepCopy() *HealthStatus {
	if in == nil {
		return nil
	}
	out := new(HealthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Operation) DeepCopyInto(out *Operation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operation.
func (in *Operation) DeepCopy() *Operation {
	if in == nil {
		return nil
	}
	out := new(Operation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Operation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationList) DeepCopyInto(out *OperationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Operation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationList.
func (in *OperationList) DeepCopy() *OperationList {
	if in == nil {
		return nil
	}
	out := new(OperationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationSpec) DeepCopyInto(out *OperationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationSpec.
func (in *OperationSpec) DeepCopy() *OperationSpec {
	if in == nil {
		return nil
	}
	out := new(OperationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationStatus) DeepCopyInto(out *OperationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationStatus.
func (in *OperationStatus) DeepCopy() *OperationStatus {
	if in == nil {
		return nil
	}
	out := new(OperationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatus) DeepCopyInto(out *ResourceStatus) {
	*out = *in
	if in.Health != nil {
		in, out := &in.Health, &out.Health
		*out = new(HealthStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatus.
func (in *ResourceStatus) DeepCopy() *ResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncStatus) DeepCopyInto(out *SyncStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncStatus.
func (in *SyncStatus) DeepCopy() *SyncStatus {
	if in == nil {
		return nil
	}
	out := new(SyncStatus)
	in.DeepCopyInto(out)
	return out
}
