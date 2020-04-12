package common

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
	"k8s.io/klog"
	"kx/kx/util"
)

const (
	defaultQPS   = 1000
	defaultBurst = 1000
)

var cf *genericclioptions.ConfigFlags

func GetConfigFlags() *genericclioptions.ConfigFlags {
	return cf
}

func apis() (*resourceLookup, error) {
	dc, err := cf.ToDiscoveryClient()
	if err != nil {
		return nil, err
	}

	resourcesList, err := dc.ServerPreferredResources()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch api groups from kubernetes: %w", err)
	}

	resourceLookup := &resourceLookup{
		resources: make(map[string][]resource),
	}

	for _, resourceList := range resourcesList {
		gv, err := schema.ParseGroupVersion(resourceList.GroupVersion)
		if err != nil {
			return nil, fmt.Errorf("%q cannot be parsed into groupversion: %w", resourceList.GroupVersion, err)
		}

		for _, api := range resourceList.APIResources {
			if !util.SliceContains(api.Verbs, "list") {
				klog.V(4).Infof("api resource \"%s\" doesn't have the verb \"list\"", api.Name, api.Verbs)
				continue
			}

			resource := resource{
				groupVersion: gv,
				apiResource:  api,
			}

			names := append([]string{api.SingularName, api.Name}, api.ShortNames...)
			for _, name := range names {
				if name == "" {
					continue
				}

				resourceLookup.resources[name] = append(resourceLookup.resources[name], resource)
			}
		}
	}

	return resourceLookup, nil
}

func list(ctx context.Context, r resource) (*unstructured.UnstructuredList, error) {
	gvr := schema.GroupVersionResource{
		Group:    r.groupVersion.Group,
		Version:  r.groupVersion.Version,
		Resource: r.apiResource.Name,
	}
	ns := GetNamespace()

	rc, err := cf.ToRESTConfig()
	if err != nil {
		return nil, err
	}

	rc.QPS = defaultQPS
	rc.Burst = defaultBurst

	dyn, err := dynamic.NewForConfig(rc)
	if err != nil {
		return nil, fmt.Errorf("failed to create dynamic client: %w", err)
	}

	var ri dynamic.ResourceInterface
	if r.apiResource.Namespaced {
		ri = dyn.Resource(gvr).Namespace(ns)
	} else {
		ri = dyn.Resource(gvr)
	}

	return ri.List(ctx, metav1.ListOptions{})
}

func init() {
	cf = genericclioptions.NewConfigFlags(true)
}
