package printer

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"strings"
)

func Print(ls *unstructured.UnstructuredList) error {
	if len(ls.Items) == 0 {
		fmt.Println("No resources found")
	} else {
		for _, item := range ls.Items {
			fmt.Print(formatMessage(item))
		}
	}

	return nil
}

func formatMessage(u unstructured.Unstructured) string {
	kind, name := strings.ToLower(u.GetKind()), u.GetName()
	return fmt.Sprintf("%s/%s\n", kind, name)
}
