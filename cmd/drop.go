package cmd

import (
	"fmt"

	"github.com/rogemus/worklog/internal/database"

	"github.com/spf13/cobra"
)

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Delete all tasks from history",
	Long:  "This command removes all previously logged tasks from your weekly history",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.NewDB(dbPath)
		err := db.RemoveAllTasks()
		if err != nil {
			fmt.Println("Error: cannot remove tasks\n", err)
			return
		}

		fmt.Printf("Success: All tasks removed\n")
	},
}
