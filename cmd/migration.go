package cmd

import (
	"github.com/s4kibs4mi/snapify/cmd/migration"
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "migration handles database migration operations",
}

func init() {
	migrationCmd.AddCommand(migration.UpCmd)
	migrationCmd.AddCommand(migration.DownCmd)
}
