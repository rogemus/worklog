package cmd

import "tempocli/internal"

var db *internal.DB

func InitCommands(database *internal.DB) {
	db = database
	rootCmd.AddCommand(findCmd)

	newCmd.Flags().StringVar(&date, "date", "", "task date in YYYY-MM-DD format (optional)")
	rootCmd.AddCommand(newCmd)

	rootCmd.AddCommand(removeCmd)

	// TODO
	// reportCmd.Flags().StringVar(&date, "date", "", "tasks from last week")
	// reportCmd.Flags().StringVar(&date, "date", "", "tasks from current month")
	// reportCmd.Flags().StringVar(&date, "date", "", "tasks from last month")
	rootCmd.AddCommand(reportCmd)
}
