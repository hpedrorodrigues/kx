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
		return printResource(resources[i])
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
		return printResource(resources[i])
	})

	if err != nil {
		return unstructured.Unstructured{}, err
	}

	return resources[idx], nil
}

func FindRolloutResources(ctx context.Context) (unstructured.Unstructured, error) {
	resources, err := ListRolloutResources(ctx)
	if err != nil {
		return unstructured.Unstructured{}, err
	}

	idx, err := fuzzyfinder.Find(resources, func(i int) string {
		return printResource(resources[i])
	})

	if err != nil {
		return unstructured.Unstructured{}, err
	}

	return resources[idx], nil
}

func printResource(u unstructured.Unstructured) string {
	kind, name := strings.ToLower(u.GetKind()), u.GetName()
	return fmt.Sprintf("%s/%s", kind, name)
}
