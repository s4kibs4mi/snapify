package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "snapify",
}

func init() {
	rootCmd.AddCommand(instanceCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
