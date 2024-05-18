package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Prints the details and status information about all the migrations",
	Long: `Prints the details and status information about all the migrations.

Info lets you know where you stand. At a glance you will see which migrations
have already been applied, which other ones are still pending, when they were
executed and whether they were successful or not.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
