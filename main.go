package main

import (
	"worklog/cmd"
	"worklog/internal/database"
)

func main() {
	db := database.NewDB(".")
	cmd.InitCommands(db)

	cmd.Execute()
}
