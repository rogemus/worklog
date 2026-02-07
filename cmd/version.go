package cmd

import (
	"fmt"

	"github.com/rogemus/worklog/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current application version",
	Long:  "This command prints the version of the application to the standard output.",
	Run: func(cmd *cobra.Command, args []string) {
		newestVersion, err := version.GetNewestVersion()
		if err != nil {
			fmt.Printf("tmxu version %s \n", appVersion)
			return
		}

		sv := version.NewSemVer(appVersion)
		if sv.Original == newestVersion.Original {
			fmt.Printf("tmxu version %s \n", appVersion)
			return
		}

		fmt.Println("A new version of the tmxu is available!")
		fmt.Println("Please run the following command to update:")
		fmt.Printf("  go install github.com/rogemus/worklog@%s\n\n", newestVersion.Original)
		fmt.Printf("Current tmxu version %s \n", appVersion)
	},
}
