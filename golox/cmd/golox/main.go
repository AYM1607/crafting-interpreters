package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/AYM1607/crafting-interpreters/golox/internal/runner"
)

func main() {
	argc := len(os.Args)
	switch {
	case argc > 2:
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	case argc == 2:
		err := runner.RunFile(os.Args[1])
		if errors.Is(err, runner.ErrInvalidScriptFile) {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		runner.RunPrompt()
	}
}
