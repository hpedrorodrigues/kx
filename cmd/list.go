package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"kx/kx"
)

var listCmd = &cobra.Command{
	Use:          "ls",
	Short:        "List resources",
	Long:         "List Kubernetes resources",
	Args:         cobra.NoArgs,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return kx.List(context.Background())
	},
}
