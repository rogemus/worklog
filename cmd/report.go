package cmd

import (
	"errors"
	"fmt"
	"time"
	"worklog/internal/database"
	"worklog/internal/reports"
	"worklog/internal/weekRange"

	"github.com/spf13/cobra"
)

// FLAGS
var (
	month bool
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Show a summary of tasks for the current week",
	Long:  "Generates a report of all tasks logged for weeks. The report can include task counts, and group tasks by day",
	Run: func(cmd *cobra.Command, args []string) {
		date := time.Now()

		var (
			rangeLabel string
			rangeTitle string
			tasksRange weekRange.WeekRange
		)

		if month {
			tasksRange = weekRange.NewMonthRange(date)
			rangeTitle = "month"
			rangeLabel = tasksRange.GetFormatedRangeForMonth()
		} else {
			tasksRange = weekRange.NewWeekRange(date)
			rangeTitle = "week"
			rangeLabel = tasksRange.GetFormatedRange()
		}

		tasks, err := db.FindTasksForWeek(tasksRange)

		if err != nil && errors.Is(err, database.ErrorNoTasks) {
			fmt.Printf("No tasks for week: %s\n", rangeLabel)
			return
		} else if err != nil {
			fmt.Println("Error: cannot find tasks", err)
			return
		}

		listTitle := fmt.Sprintf("\nTasks [#%d] for %s: %s", len(tasks), rangeTitle, rangeLabel)
		fmt.Println(listTitle)

		reports.RenderAsList(tasks)
	},
}
