package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from your history",
	Long:  "This command removes a previously logged task from your weekly history",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}
