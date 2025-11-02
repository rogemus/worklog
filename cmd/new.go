package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
	"worklog/internal"

	"github.com/spf13/cobra"
)

var (
	date  string
	isGit bool
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new task to your weekly log",
	Long:  "Use this command to record a new task you've worked on. You can include details such as the task name, description, date",
	Run: func(cmd *cobra.Command, args []string) {
		var taskName string
		taskDate := time.Now()

		if len(args) > 0 {
			taskName = strings.Join(args, " ")
		}

		if taskName == "" && !isGit {
			fmt.Printf("%s\n", internal.ErrorEmptyName.Error())
			return
		}

		if date != "" {
			parsedDate, err := time.Parse("2006-01-02", date)

			if err != nil {
				fmt.Printf("%s\n", internal.ErrorInvalidDate)
				return
			}

			taskDate = parsedDate
		}

		if isGit {
			cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
			cmdOutput, err := cmd.Output()

			if err != nil {
				fmt.Printf("%s\n", err.Error())
				return
			}

			if strings.TrimSpace(string(cmdOutput)) != "true" {
				fmt.Printf("%s\n", internal.ErrorNotInGitRepo.Error())
				return
			}

			cmd = exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
			cmdOutput, err = cmd.Output()
			branchName := strings.TrimSpace(string(cmdOutput))

			if err != nil {
				fmt.Printf("%s\n", internal.ErrorNotInGitRepo.Error())
				return
			}

			if branchName == "main" || branchName == "master" {
				fmt.Printf("%s\n", internal.ErrorMainBranch.Error())
				return
			}

			taskName = branchName
		}

		newTask, err := db.AddTask(internal.NewTask(taskName, taskDate, nil))
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		fmt.Printf("Success: Task created id [%d] [%s]\n", newTask.Id, newTask.Name)
	},
}
