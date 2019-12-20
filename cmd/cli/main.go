package main

import "github.com/dmytro-kolesnyk/dds/cmd/cli/app"

func main() {
	cliApp := app.NewApp("localhost", "8081")
	cliApp.Start()
}