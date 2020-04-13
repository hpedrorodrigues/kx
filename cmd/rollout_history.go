package cmd

import (
	"context"
	"github.com/spf13/cobra"
	kxcmd "kx/kx/cmd"
)

var rolloutHistoryCmd = &cobra.Command{
	Use:          "rh",
	Short:        "View previous rollout revisions and configurations",
	Long:         "View previous rollout revisions and configurations",
	Args:         cobra.NoArgs,
	SuggestFor:   []string{"rollout", "rollout-history"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return kxcmd.Run(kxcmd.NewRolloutHistoryOptions(context.TODO(), cmd))
	},
}
