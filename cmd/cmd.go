package cmd

var dbPath string
var appVersion string

func InitCommands(version string) {
	appVersion = version
	rootCmd.PersistentFlags().StringVarP(&dbPath, "db-path", "p", ".config/worklog", "location of the database file")

	// VERSION
	rootCmd.AddCommand(versionCmd)

	// FIND
	rootCmd.AddCommand(findCmd)

	// NEW
	newCmd.Flags().StringVarP(&newDate, "date", "d", "", "task date in YYYY-MM-DD format (optional)")
	newCmd.Flags().BoolVarP(&isGit, "git", "g", false, "use git branch as task name")
	rootCmd.AddCommand(newCmd)

	// REMOVE
	rootCmd.AddCommand(removeCmd)

	// DROP
	rootCmd.AddCommand(dropCmd)

	// REPORT
	reportCmd.Flags().BoolVarP(&month, "month", "m", false, "report tasks from whole month")
	reportCmd.Flags().StringVarP(&reportDate, "date", "d", "", "task date in YYYY-MM-DD format (optional)")
	rootCmd.AddCommand(reportCmd)
}
