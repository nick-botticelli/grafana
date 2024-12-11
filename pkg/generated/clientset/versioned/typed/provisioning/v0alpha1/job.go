// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by client-gen. DO NOT EDIT.

package v0alpha1

import (
	"context"

	v0alpha1 "github.com/grafana/grafana/pkg/apis/provisioning/v0alpha1"
	provisioningv0alpha1 "github.com/grafana/grafana/pkg/generated/applyconfiguration/provisioning/v0alpha1"
	scheme "github.com/grafana/grafana/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// JobsGetter has a method to return a JobInterface.
// A group's client should implement this interface.
type JobsGetter interface {
	Jobs(namespace string) JobInterface
}

// JobInterface has methods to work with Job resources.
type JobInterface interface {
	Create(ctx context.Context, job *v0alpha1.Job, opts v1.CreateOptions) (*v0alpha1.Job, error)
	Update(ctx context.Context, job *v0alpha1.Job, opts v1.UpdateOptions) (*v0alpha1.Job, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, job *v0alpha1.Job, opts v1.UpdateOptions) (*v0alpha1.Job, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v0alpha1.Job, error)
	List(ctx context.Context, opts v1.ListOptions) (*v0alpha1.JobList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v0alpha1.Job, err error)
	Apply(ctx context.Context, job *provisioningv0alpha1.JobApplyConfiguration, opts v1.ApplyOptions) (result *v0alpha1.Job, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, job *provisioningv0alpha1.JobApplyConfiguration, opts v1.ApplyOptions) (result *v0alpha1.Job, err error)
	JobExpansion
}

// jobs implements JobInterface
type jobs struct {
	*gentype.ClientWithListAndApply[*v0alpha1.Job, *v0alpha1.JobList, *provisioningv0alpha1.JobApplyConfiguration]
}

// newJobs returns a Jobs
func newJobs(c *ProvisioningV0alpha1Client, namespace string) *jobs {
	return &jobs{
		gentype.NewClientWithListAndApply[*v0alpha1.Job, *v0alpha1.JobList, *provisioningv0alpha1.JobApplyConfiguration](
			"jobs",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v0alpha1.Job { return &v0alpha1.Job{} },
			func() *v0alpha1.JobList { return &v0alpha1.JobList{} }),
	}
}
