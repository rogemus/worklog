package cmd

import "worklog/internal"

var db *internal.DB

func InitCommands(database *internal.DB) {
	db = database
	// FIND
	rootCmd.AddCommand(findCmd)

	// NEW
	newCmd.Flags().StringVar(&date, "date", "", "task date in YYYY-MM-DD format (optional)")
	newCmd.Flags().BoolVar(&isGit, "git", false, "use git branch as task name")
	rootCmd.AddCommand(newCmd)

	// REMOVE
	rootCmd.AddCommand(removeCmd)

	// DROP
	rootCmd.AddCommand(dropCmd)

	// REPORT
	rootCmd.AddCommand(reportCmd)
}
