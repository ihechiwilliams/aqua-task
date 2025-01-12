package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aqua-task",
	Short: "A CLI for managing aqua-task operations",
	Long:  `A CLI tool to manage database migrations, seeding, and other operations for the aqua-task project.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
