package main

import (
	"k8s.io/klog"
	"kx/cmd"
	"os"
)

func main() {
	defer klog.Flush()

	if err := cmd.Execute(); err != nil {
		klog.Errorln(err)
		os.Exit(1)
	}
}
