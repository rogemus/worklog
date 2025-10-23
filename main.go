package main

import (
	"tempocli/cmd"
	"tempocli/internal"
)

type Config struct {
	DB internal.DB
}

func main() {
	db := internal.NewDb()
	cmd.InitCommands(db)

	cmd.Execute()
}
