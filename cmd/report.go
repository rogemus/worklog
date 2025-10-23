package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Show a summary of tasks for the current week",
	Long:  "Generates a report of all tasks logged for week. The report can include task counts, and group tasks by day",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("report called")
	},
}

