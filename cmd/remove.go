package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Delete a task from your history",
	Long:  "This command removes a previously logged task from your weekly history",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}
