package cmd

import (
	"context"
	"github.com/spf13/cobra"
	kxcmd "kx/kx/cmd"
)

var editCmd = &cobra.Command{
	Use:          "ed",
	Short:        "Edit resources",
	Long:         "Edit Kubernetes resources",
	Args:         cobra.NoArgs,
	SuggestFor:   []string{"edit"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return kxcmd.Run(kxcmd.NewEditOptions(context.TODO(), cmd))
	},
}
