package cmd

import "tempocli/internal"

var db *internal.DB

func InitCommands(database *internal.DB) {
	db = database
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(findCmd)

	newCmd.Flags().StringVar(&date, "date", "", "task date in YYYY-MM-DD format (optional)")
	rootCmd.AddCommand(newCmd)

	rootCmd.AddCommand(removeCmd)

	rootCmd.AddCommand(reportCmd)
}
