package cmd

import "worklog/internal/database"

var db *database.DB

func InitCommands(datab *database.DB) {
	db = datab
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
