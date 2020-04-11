package common

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"strings"
)

type Resource struct {
	GroupVersion schema.GroupVersion
	APIResource  metav1.APIResource
}

type ResourceLookup struct {
	resources map[string][]Resource
}

func (rl ResourceLookup) Lookup(k string) ([]Resource, bool) {
	v, ok := rl.resources[strings.ToLower(k)]
	return v, ok
}

func (rl ResourceLookup) LookupFirst(k string) (Resource, bool) {
	resources, ok := rl.Lookup(k)

	if !ok || len(resources) == 0 {
		return Resource{}, false
	}

	return resources[0], true
}
