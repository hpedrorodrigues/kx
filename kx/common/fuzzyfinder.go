package common

import (
	"context"
	"fmt"
	"github.com/ktr0731/go-fuzzyfinder"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"strings"
)

func FindResource(ctx context.Context) (unstructured.Unstructured, error) {
	resources, err := ListResources(ctx)
	if err != nil {
		return unstructured.Unstructured{}, err
	}

	idx, err := fuzzyfinder.Find(resources, func(i int) string {
		r := resources[i]
		kind, name := strings.ToLower(r.GetKind()), r.GetName()
		return fmt.Sprintf("%s/%s", kind, name)
	})

	if err != nil {
		return unstructured.Unstructured{}, err
	}

	return resources[idx], nil
}

func FindPod(ctx context.Context) (unstructured.Unstructured, error) {
	resources, err := ListPods(ctx)
	if err != nil {
		return unstructured.Unstructured{}, err
	}

	idx, err := fuzzyfinder.Find(resources, func(i int) string {
		r := resources[i]
		kind, name := strings.ToLower(r.GetKind()), r.GetName()
		return fmt.Sprintf("%s/%s", kind, name)
	})

	if err != nil {
		return unstructured.Unstructured{}, err
	}

	return resources[idx], nil
}

func FindDDS(ctx context.Context) (unstructured.Unstructured, error) {
	resources, err := ListDDS(ctx)
	if err != nil {
		return unstructured.Unstructured{}, err
	}

	idx, err := fuzzyfinder.Find(resources, func(i int) string {
		r := resources[i]
		kind, name := strings.ToLower(r.GetKind()), r.GetName()
		return fmt.Sprintf("%s/%s", kind, name)
	})

	if err != nil {
		return unstructured.Unstructured{}, err
	}

	return resources[idx], nil
}
