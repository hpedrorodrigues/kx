package cmd

import (
	"context"
	"github.com/spf13/cobra"
	kxcmd "kx/kx/cmd"
)

var rolloutStatusCmd = &cobra.Command{
	Use:          "rs",
	Short:        "Show the status of the rollout",
	Long:         "Show the status of the rollout",
	Args:         cobra.NoArgs,
	SuggestFor:   []string{"rollout", "rollout-status"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return kxcmd.Run(kxcmd.NewRolloutStatusOptions(context.TODO(), cmd))
	},
}
