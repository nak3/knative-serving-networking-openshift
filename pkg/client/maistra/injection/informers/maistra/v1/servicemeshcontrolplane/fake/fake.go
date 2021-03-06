// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	"context"

	fake "github.com/openshift-knative/knative-serving-networking-openshift/pkg/client/maistra/injection/informers/factory/fake"
	servicemeshcontrolplane "github.com/openshift-knative/knative-serving-networking-openshift/pkg/client/maistra/injection/informers/maistra/v1/servicemeshcontrolplane"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = servicemeshcontrolplane.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Maistra().V1().ServiceMeshControlPlanes()
	return context.WithValue(ctx, servicemeshcontrolplane.Key{}, inf), inf.Informer()
}
