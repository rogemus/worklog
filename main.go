package main

import (
	"worklog/cmd"
)

func main() {
	cmd.InitCommands()
	cmd.Execute()
}
