package cmd

import (
	"context"
	"github.com/spf13/cobra"
	kxcmd "kx/kx/cmd"
)

var describeCmd = &cobra.Command{
	Use:          "de",
	Short:        "Describe resources",
	Long:         "Print information about Kubernetes resources",
	Args:         cobra.NoArgs,
	SuggestFor:   []string{"describe"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return kxcmd.Run(kxcmd.NewDescribeOptions(context.TODO(), cmd))
	},
}
