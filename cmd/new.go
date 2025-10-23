package cmd

import (
	"fmt"
	"strings"
	"tempocli/internal"
	"time"

	"github.com/spf13/cobra"
)

var (
	date string
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new task to your weekly log",
	Long:  "Use this command to record a new task you've worked on. You can include details such as the task name, description, date",
	Run: func(cmd *cobra.Command, args []string) {
		var taskName string
		taskDate := time.Now()

		// Get input text from args or stdin
		if len(args) > 0 {
			taskName = strings.Join(args, " ")
		}

		if taskName == "" {
			fmt.Println("No input text provided")
			return
		}

		if date != "" {
			parsedDate, err := time.Parse("2006-01-02", date)

			if err != nil {
				fmt.Println("Invalid task date format, using current day as task date")
				return
			}

			taskDate = parsedDate
		}

		newTask := internal.NewTask(taskName, taskDate, nil)
		_, err := db.AddTask(newTask)
		if err != nil {
			panic(err)
		}
		fmt.Println("task created")
	},
}
