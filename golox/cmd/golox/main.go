package main

import (
	"fmt"
	"os"

	"github.com/AYM1607/crafting-interpreters/golox/internal/runner"
)

func main() {
	argc := len(os.Args)
	if argc > 2 {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	} else if argc == 2 {

	} else {
		runner.RunPrompt()
	}
}
