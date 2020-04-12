package cmd

import (
	"github.com/spf13/cobra"
	"kx/kx"
)

const sortByFlag = "sort-by"

var listCmd = &cobra.Command{
	Use:          "ls",
	Short:        "List resources",
	Long:         "List Kubernetes resources",
	Args:         cobra.NoArgs,
	SuggestFor:   []string{"list"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return kx.List(cmd)
	},
}

func init() {
	listCmd.Flags().String(
		sortByFlag,
		"",
		"If non-empty, sort list types using this field specification (e.g. '{.metadata.name}').",
	)
}
