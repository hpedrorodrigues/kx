package common

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"kx/kx/util"
	"os"
	"strings"
)

const (
	resourceEnvName  = "KX_RESOURCES"
	defaultResources = "pods,deployments"
)

var iostreams genericclioptions.IOStreams

func GetIOStreams() genericclioptions.IOStreams {
	return iostreams
}

func ListResources(ctx context.Context) ([]unstructured.Unstructured, error) {
	apis, err := apis()
	if err != nil {
		return nil, err
	}

	var items []unstructured.Unstructured

	names := strings.Split(GetResourceNames(), ",")
	for _, name := range names {
		r, ok := apis.LookupFirst(name)
		if !ok {
			fmt.Printf("invalid resource: %s\n", name)
			continue
		}

		list, err := list(ctx, r)
		if err != nil {
			return nil, err
		}

		items = append(items, list.Items...)
	}

	return items, nil
}

func ListPods(ctx context.Context) ([]unstructured.Unstructured, error) {
	apis, err := apis()
	if err != nil {
		return nil, err
	}

	name := "pods"
	r, ok := apis.LookupFirst(name)
	if !ok {
		fmt.Printf("invalid resource: %s\n", name)
	}

	list, err := list(ctx, r)
	if err != nil {
		return nil, err
	}

	return list.Items, nil
}

func ListRolloutResources(ctx context.Context) ([]unstructured.Unstructured, error) {
	apis, err := apis()
	if err != nil {
		return nil, err
	}

	var items []unstructured.Unstructured
	filters := []string{"deployments", "daemonsets", "statefulsets"}

	names := strings.Split(GetResourceNames(), ",")
	for _, name := range names {
		r, ok := apis.LookupFirst(name)
		if !ok {
			fmt.Printf("invalid resource: %s\n", name)
			continue
		}

		if !util.SliceContains(filters, name) {
			continue
		}

		list, err := list(ctx, r)
		if err != nil {
			return nil, err
		}

		items = append(items, list.Items...)
	}

	return items, nil
}

func GetResourceNames() string {
	res := os.Getenv(resourceEnvName)
	if res == "" {
		res = defaultResources
	}

	return res
}

func init() {
	iostreams = genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
}
