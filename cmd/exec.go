package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"kx/kx"
	"time"
)

const (
	podRunningTimeoutFlag = "pod-running-timeout"
	defaultPodExecTimeout = 60 * time.Second
)

var execCmd = &cobra.Command{
	Use:          "ex",
	Short:        "Execute commands inside a container in a pod",
	Long:         "Execute commands inside a container in a pod",
	Args:         cobra.NoArgs,
	SuggestFor:   []string{"exec"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return kx.Exec(context.TODO(), cmd)
	},
}

func init() {
	execCmd.Flags().Duration(
		podRunningTimeoutFlag,
		defaultPodExecTimeout,
		"The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running",
	)
}
