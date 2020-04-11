package kx

import (
	"context"
	"fmt"
	"kx/kx/common"
	"kx/kx/printer"
	"os"
	"strings"
)

const (
	resourceEnvName  = "KX_RESOURCES"
	defaultResources = "pods,deployments"
)

func List(ctx context.Context) error {
	apis, err := common.APIs()
	if err != nil {
		return err
	}

	for _, name := range getResourceNames() {
		r, ok := apis.LookupFirst(name)
		if !ok {
			fmt.Printf("invalid resource: %s\n", name)
			continue
		}

		list, err := common.List(ctx, r)
		if err != nil {
			return err
		}

		if err := printer.Print(list); err != nil {
			return err
		}
	}

	return nil
}

func getResourceNames() []string {
	res := os.Getenv(resourceEnvName)
	if res == "" {
		res = defaultResources
	}

	return strings.Split(res, ",")
}
