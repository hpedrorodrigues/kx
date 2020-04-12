package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"kx/kx"
)

var rolloutStatusCmd = &cobra.Command{
	Use:          "rs",
	Short:        "Show the status of the rollout",
	Long:         "Show the status of the rollout",
	Args:         cobra.NoArgs,
	SuggestFor:   []string{"rollout", "rollout-status"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return kx.RolloutStatus(context.TODO())
	},
}
