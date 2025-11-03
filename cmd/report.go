package cmd

import (
	"errors"
	"fmt"
	"time"
	"worklog/internal"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Show a summary of tasks for the current week",
	Long:  "Generates a report of all tasks logged for week. The report can include task counts, and group tasks by day",
	Run: func(cmd *cobra.Command, args []string) {
		date := time.Now()

		week := internal.NewWeekRange(date)
		tasks, err := db.FindTasksForWeek(week)

		if err != nil && errors.Is(err, internal.ErrorNoTasks) {
			fmt.Printf("No tasks for week: [%s - %s]\n", week.Start, week.End)
			return
		} else if err != nil {
			fmt.Println("Error: cannot find tasks", err)
			return
		}

		listTitle := fmt.Sprintf("Tasks for week: %s", week.GetFormatedRange())
		fmt.Println(listTitle)

		internal.RenderAsList(tasks)
	},
}
