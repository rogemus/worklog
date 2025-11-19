package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
	"worklog/internal/database"
	"worklog/internal/task"

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
		db := database.NewDB(dbPath)

		taskDate := time.Now()
		taskName := strings.TrimSpace(strings.Join(args, " "))

		if date != "" {
			parsedDate, err := time.Parse("2006-01-02", date)

			if err != nil {
				fmt.Printf("%s\n", task.ErrorParseInvalidCreatedDate.Error())
				return
			}

			taskDate = parsedDate
		}

		if isGit {
			if !isInGitRepo() {
				fmt.Printf("%s\n", task.ErrorNotInGitRepo.Error())
				return
			}

			repoName := getGitRepoName()
			branch := getGitBranch()

			newTask, err := db.AddTask(task.NewTaskFromGit(branch, repoName, taskDate))

			if err != nil {
				fmt.Printf("%s\n", err.Error())
				return
			}

			fmt.Printf("Success: Task created id [%s] [%s]\n", newTask.GetFormattedId(), newTask.GetFormattedName())
			return
		}

		if taskName == "" {
			fmt.Printf("%s\n", task.ErrorEmptyName.Error())
			return
		}

		newTask, err := db.AddTask(task.NewTask(taskName, taskDate, nil))
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		fmt.Printf("Success: Task created id [%s] [%s]\n", newTask.GetFormattedId(), newTask.GetFormattedName())
	},
}

func isInGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmdOutput, err := cmd.Output()

	if err != nil {
		return false
	}

	outputStr := strings.TrimSpace(string(cmdOutput))
	if outputStr != "true" {
		return false
	}

	return true
}

func getGitRepoName() string {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	cmdOutput, err := cmd.Output()

	if err != nil {
		return ""
	}

	origin := strings.TrimSpace(string(cmdOutput))
	parts := strings.Split(origin, "/")

	return strings.ReplaceAll(parts[len(parts)-1], ".git", "")
}

func getGitBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmdOutput, err := cmd.Output()

	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(cmdOutput))
}
