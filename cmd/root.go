package cmd

import "github.com/spf13/cobra"

const version = "v0.0.1"

var rootCmd = &cobra.Command{
	Use:     "kx",
	Short:   "A utility command-line tool to improve productivity with day to day tasks",
	Long:    "A utility command-line tool to improve productivity with day to day tasks dealing with Kubernetes resources",
	Version: version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(describeCmd)
	rootCmd.AddCommand(listCmd)
}
