package common

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"strings"
)

type resource struct {
	groupVersion schema.GroupVersion
	apiResource  metav1.APIResource
}

type resourceLookup struct {
	resources map[string][]resource
}

func (rl resourceLookup) Lookup(k string) ([]resource, bool) {
	v, ok := rl.resources[strings.ToLower(k)]
	return v, ok
}

func (rl resourceLookup) LookupFirst(k string) (resource, bool) {
	resources, ok := rl.Lookup(k)

	if !ok || len(resources) == 0 {
		return resource{}, false
	}

	return resources[0], true
}
