package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "worklog",
	Short: "Log and manage weekly tasks from the command line",
	Long: `worklog is a lightweight CLI tool for tracking what you work on during the week.
	It lets you quickly add, edit, remove, and report tasks, helping you stay organized and reflect on your weekly progress.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
