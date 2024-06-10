package main

import (
	"fmt"
	"os"

	"pubbo/cmd"
)

func main () {
	os.Exit(Run())
}

func Run () int {
	rootCmd := cmd.CreateRootCommand()
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}

