package main

import (
	"fmt"
	"k8s.io/klog"
	"kx/cmd"
	"os"
)

func main() {
	defer klog.Flush()

	if err := cmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
