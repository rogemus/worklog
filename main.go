package main

import (
	_ "embed"
	"github.com/rogemus/worklog/cmd"
)

//go:embed version.txt
var version string

func main() {
	cmd.InitCommands(version)
	cmd.Execute()
}
