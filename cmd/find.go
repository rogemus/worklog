package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rogemus/worklog/internal/database"
	"github.com/rogemus/worklog/internal/reports"
	"github.com/rogemus/worklog/internal/task"

	"github.com/spf13/cobra"
)

var (
	// TODO: not implemented
	findByTags       bool
	findByRepoName   bool
	findByBranchName bool
	findByIssue      bool
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

		var (
			tasks     []task.Task
			err       error
			query     = strings.Join(args, " ")
			db        = database.NewDB(dbPath)
			modifiers = []bool{
				findByTags,
				findByRepoName,
				findByBranchName,
				findByIssue,
			}
		)

		if hasMultipleFlagSet(modifiers) {
			fmt.Printf("%s\n", task.ErrorFindMultipleFlags.Error())
			return
		}

		// TODO: (for future relase) allow to use all flags at once

		switch {
		case findByBranchName:
			tasks, err = db.FindTasks(query, "branch")
		case findByIssue:
			tasks, err = db.FindTasks(query, "issue")
		case findByRepoName:
			tasks, err = db.FindTasks(query, "repo")
		case findByTags:
			tasks, err = db.FindTasksWithTags(query)
		default:
			tasks, err = db.FindTasks(query, "name")
		}

		if err != nil && errors.Is(err, database.ErrorNoTasks) {
			fmt.Printf("No tasks found for query: %s\n", query)
			return
		} else if err != nil {
			fmt.Println("Error: cannot find tasks", err)
			return
		}

		listTitle := fmt.Sprintf("\nTasks [#%d] for query [%s]", len(tasks), query)
		fmt.Println(listTitle)
		reports.RenderAsList(tasks)
	},
}

func hasMultipleFlagSet(modifiers []bool) bool {
	found := []bool{}

	for _, modifier := range modifiers {
		if modifier {

			if len(found) == 1 {
				// NOTE: one true flag is saved

				return true
			}

			found = append(found, modifier)
		}
	}

	return false
}
