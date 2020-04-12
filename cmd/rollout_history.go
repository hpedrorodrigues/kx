package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"kx/kx"
)

var rolloutHistoryCmd = &cobra.Command{
	Use:          "rh",
	Short:        "View previous rollout revisions and configurations",
	Long:         "View previous rollout revisions and configurations",
	Args:         cobra.NoArgs,
	SuggestFor:   []string{"rollout", "rollout-history"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return kx.RolloutHistory(context.TODO(), cmd)
	},
}
