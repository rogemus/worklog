package cmd

import (
	"errors"
	"fmt"
	"strings"
	"tempocli/internal"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Search for tasks in your history.",
	Long:  "Use this command to search for tasks you've logged by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: Provide query")
			return
		}

		query := strings.Join(args, " ")
		tasks, err := db.FindTasksByName(query)

		if err != nil && errors.Is(err, internal.ErrorNoTasks) {
			fmt.Printf("No tasks found for query: %s\n", query)
			return
		} else if err != nil {
			fmt.Println("Error: cannot find tasks", err)
			return
		}

		for _, task := range tasks {
			fmt.Printf("%s [%d] %s\n", task.Date, task.Id, task.Name)
		}
	},
}
