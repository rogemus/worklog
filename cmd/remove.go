package cmd

import (
	"fmt"
	"strconv"
	"worklog/internal/database"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Delete a task from your history",
	Long:  "This command removes a previously logged task from your weekly history",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Error: Provide single task id")
			return
		}

		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: invalid task id")
			return
		}

		db := database.NewDB(dbPath)
		err = db.RemoveTask(taskId)
		if err != nil {
			fmt.Println("Error: cannot remove task\n", err)
			return
		}

		fmt.Printf("Success: Task [%d] removed\n", taskId)
	},
}
