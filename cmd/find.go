package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Search for tasks in your history.",
	Long:  "Use this command to search for tasks you've logged by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("find called")
	},
}
