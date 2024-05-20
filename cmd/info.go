package cmd

import (
	"github.com/eidolons/nyx/extractors"
	"github.com/eidolons/nyx/services"
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
		data := extractors.InfoSettings(*cmd, args)

		services.Run(data)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
