package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current application version",
	Long:  "This command prints the version of the application to the standard output.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s", appVersion)
	},
}
