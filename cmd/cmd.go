package cmd

import "tempocli/internal"

var db *internal.DB

func InitCommands(database *internal.DB) {
	db = database
	// FIND
	rootCmd.AddCommand(findCmd)

	// NEW
	newCmd.Flags().StringVar(&date, "date", "", "task date in YYYY-MM-DD format (optional)")
	rootCmd.AddCommand(newCmd)

	// REMOVE
	rootCmd.AddCommand(removeCmd)

	// DROP
	rootCmd.AddCommand(dropCmd)

	// REPORT
	reportCmd.Flags().BoolVar(&table, "table", false, "render report as table")
	rootCmd.AddCommand(reportCmd)
}
