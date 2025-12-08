package main

import (
	"github.com/rogemus/worklog/cmd"
)

func main() {
	cmd.InitCommands()
	cmd.Execute()
}
