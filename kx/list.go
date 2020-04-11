package kx

import (
	"context"
	"fmt"
	"kx/kx/common"
)

func List(ctx context.Context) error {
	apis, err := common.APIs()
	if err != nil {
		return err
	}

	r := apis.Resources["pods"]
	if err := printResources(ctx, r); err != nil {
		return err
	}

	return nil
}

func printResources(ctx context.Context, r common.Resource) error {
	list, err := common.List(ctx, r)
	if err != nil {
		return err
	}

	if len(list.Items) == 0 {
		fmt.Println("No resources found")
	} else {
		for _, item := range list.Items {
			fmt.Printf("%s/%s\n", item.GetKind(), item.GetName())
		}
	}

	return nil
}
