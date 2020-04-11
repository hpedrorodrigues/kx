package common

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type Resource struct {
	GroupVersion schema.GroupVersion
	APIResource  metav1.APIResource
}

type ResourceLookup struct {
	Resources map[string]Resource
}
