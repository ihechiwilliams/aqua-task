package cmd

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the application without seeding the database",
	Long:  `This command starts the application without performing any database seeding.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	// Register the startCmd command
	rootCmd.AddCommand(startCmd)
}
