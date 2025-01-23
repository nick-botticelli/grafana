// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v0alpha1 "github.com/grafana/grafana/pkg/apis/service/v0alpha1"
	servicev0alpha1 "github.com/grafana/grafana/pkg/generated/applyconfiguration/service/v0alpha1"
	typedservicev0alpha1 "github.com/grafana/grafana/pkg/generated/clientset/versioned/typed/service/v0alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeExternalNames implements ExternalNameInterface
type fakeExternalNames struct {
	*gentype.FakeClientWithListAndApply[*v0alpha1.ExternalName, *v0alpha1.ExternalNameList, *servicev0alpha1.ExternalNameApplyConfiguration]
	Fake *FakeServiceV0alpha1
}

func newFakeExternalNames(fake *FakeServiceV0alpha1, namespace string) typedservicev0alpha1.ExternalNameInterface {
	return &fakeExternalNames{
		gentype.NewFakeClientWithListAndApply[*v0alpha1.ExternalName, *v0alpha1.ExternalNameList, *servicev0alpha1.ExternalNameApplyConfiguration](
			fake.Fake,
			namespace,
			v0alpha1.SchemeGroupVersion.WithResource("externalnames"),
			v0alpha1.SchemeGroupVersion.WithKind("ExternalName"),
			func() *v0alpha1.ExternalName { return &v0alpha1.ExternalName{} },
			func() *v0alpha1.ExternalNameList { return &v0alpha1.ExternalNameList{} },
			func(dst, src *v0alpha1.ExternalNameList) { dst.ListMeta = src.ListMeta },
			func(list *v0alpha1.ExternalNameList) []*v0alpha1.ExternalName {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v0alpha1.ExternalNameList, items []*v0alpha1.ExternalName) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
