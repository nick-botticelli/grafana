//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by deepcopy-gen. DO NOT EDIT.

package v0alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Author) DeepCopyInto(out *Author) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Author.
func (in *Author) DeepCopy() *Author {
	if in == nil {
		return nil
	}
	out := new(Author)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EditingOptions) DeepCopyInto(out *EditingOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EditingOptions.
func (in *EditingOptions) DeepCopy() *EditingOptions {
	if in == nil {
		return nil
	}
	out := new(EditingOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileItem) DeepCopyInto(out *FileItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileItem.
func (in *FileItem) DeepCopy() *FileItem {
	if in == nil {
		return nil
	}
	out := new(FileItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileList) DeepCopyInto(out *FileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FileItem, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileList.
func (in *FileList) DeepCopy() *FileList {
	if in == nil {
		return nil
	}
	out := new(FileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileRef) DeepCopyInto(out *FileRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileRef.
func (in *FileRef) DeepCopy() *FileRef {
	if in == nil {
		return nil
	}
	out := new(FileRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubRepositoryConfig) DeepCopyInto(out *GitHubRepositoryConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubRepositoryConfig.
func (in *GitHubRepositoryConfig) DeepCopy() *GitHubRepositoryConfig {
	if in == nil {
		return nil
	}
	out := new(GitHubRepositoryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelloWorld) DeepCopyInto(out *HelloWorld) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelloWorld.
func (in *HelloWorld) DeepCopy() *HelloWorld {
	if in == nil {
		return nil
	}
	out := new(HelloWorld)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelloWorld) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistoryItem) DeepCopyInto(out *HistoryItem) {
	*out = *in
	if in.Authors != nil {
		in, out := &in.Authors, &out.Authors
		*out = make([]Author, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistoryItem.
func (in *HistoryItem) DeepCopy() *HistoryItem {
	if in == nil {
		return nil
	}
	out := new(HistoryItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistoryList) DeepCopyInto(out *HistoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HistoryItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistoryList.
func (in *HistoryList) DeepCopy() *HistoryList {
	if in == nil {
		return nil
	}
	out := new(HistoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HistoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Job) DeepCopyInto(out *Job) {
	*out = *in
	if in.Added != nil {
		in, out := &in.Added, &out.Added
		*out = make([]FileRef, len(*in))
		copy(*out, *in)
	}
	if in.Modified != nil {
		in, out := &in.Modified, &out.Modified
		*out = make([]FileRef, len(*in))
		copy(*out, *in)
	}
	if in.Removed != nil {
		in, out := &in.Removed, &out.Removed
		*out = make([]FileRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Job.
func (in *Job) DeepCopy() *Job {
	if in == nil {
		return nil
	}
	out := new(Job)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LintIssue) DeepCopyInto(out *LintIssue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LintIssue.
func (in *LintIssue) DeepCopy() *LintIssue {
	if in == nil {
		return nil
	}
	out := new(LintIssue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalRepositoryConfig) DeepCopyInto(out *LocalRepositoryConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRepositoryConfig.
func (in *LocalRepositoryConfig) DeepCopy() *LocalRepositoryConfig {
	if in == nil {
		return nil
	}
	out := new(LocalRepositoryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Repository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryList) DeepCopyInto(out *RepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryList.
func (in *RepositoryList) DeepCopy() *RepositoryList {
	if in == nil {
		return nil
	}
	out := new(RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpec) DeepCopyInto(out *RepositorySpec) {
	*out = *in
	out.Editing = in.Editing
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalRepositoryConfig)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3RepositoryConfig)
		**out = **in
	}
	if in.GitHub != nil {
		in, out := &in.GitHub, &out.GitHub
		*out = new(GitHubRepositoryConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpec.
func (in *RepositorySpec) DeepCopy() *RepositorySpec {
	if in == nil {
		return nil
	}
	out := new(RepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryStatus) DeepCopyInto(out *RepositoryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryStatus.
func (in *RepositoryStatus) DeepCopy() *RepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceObjects) DeepCopyInto(out *ResourceObjects) {
	*out = *in
	out.Type = in.Type
	in.File.DeepCopyInto(&out.File)
	in.Existing.DeepCopyInto(&out.Existing)
	in.DryRun.DeepCopyInto(&out.DryRun)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceObjects.
func (in *ResourceObjects) DeepCopy() *ResourceObjects {
	if in == nil {
		return nil
	}
	out := new(ResourceObjects)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceType) DeepCopyInto(out *ResourceType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceType.
func (in *ResourceType) DeepCopy() *ResourceType {
	if in == nil {
		return nil
	}
	out := new(ResourceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceWrapper) DeepCopyInto(out *ResourceWrapper) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = (*in).DeepCopy()
	}
	in.Resource.DeepCopyInto(&out.Resource)
	if in.Lint != nil {
		in, out := &in.Lint, &out.Lint
		*out = make([]LintIssue, len(*in))
		copy(*out, *in)
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceWrapper.
func (in *ResourceWrapper) DeepCopy() *ResourceWrapper {
	if in == nil {
		return nil
	}
	out := new(ResourceWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceWrapper) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3RepositoryConfig) DeepCopyInto(out *S3RepositoryConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3RepositoryConfig.
func (in *S3RepositoryConfig) DeepCopy() *S3RepositoryConfig {
	if in == nil {
		return nil
	}
	out := new(S3RepositoryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookResponse) DeepCopyInto(out *WebhookResponse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Jobs != nil {
		in, out := &in.Jobs, &out.Jobs
		*out = make([]Job, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookResponse.
func (in *WebhookResponse) DeepCopy() *WebhookResponse {
	if in == nil {
		return nil
	}
	out := new(WebhookResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebhookResponse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
