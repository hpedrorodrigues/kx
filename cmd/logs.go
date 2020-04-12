package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"kx/kx"
)

const (
	tailFlag      = "tail"
	containerFlag = "container"
)

var logsCmd = &cobra.Command{
	Use:          "lo",
	Short:        "Print logs of a container in a pod",
	Long:         "Print logs of a container in a pod",
	Args:         cobra.NoArgs,
	SuggestFor:   []string{"logs", "log"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return kx.Logs(context.TODO(), cmd)
	},
}

func init() {
	logsCmd.Flags().Int(
		tailFlag,
		0,
		"Line numbers",
	)

	logsCmd.Flags().String(
		containerFlag,
		"",
		"Container name",
	)

	logsCmd.Flags().Duration(
		podRunningTimeoutFlag,
		defaultPodExecTimeout,
		"The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running",
	)
}
