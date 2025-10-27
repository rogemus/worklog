package main

import (
	"worklog/cmd"
	"worklog/internal"
)

type Config struct {
	DB internal.DB
}

func main() {
	db := internal.NewDb()
	cmd.InitCommands(db)

	cmd.Execute()
}
