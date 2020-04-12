package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"kx/kx"
)

var editCmd = &cobra.Command{
	Use:          "ed",
	Short:        "Edit resources",
	Long:         "Edit Kubernetes resources",
	Args:         cobra.NoArgs,
	SuggestFor:   []string{"edit"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return kx.Edit(context.TODO(), cmd)
	},
}
