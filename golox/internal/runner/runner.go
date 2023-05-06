package runner

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	lerrors "github.com/AYM1607/crafting-interpreters/golox/internal/errors"
)

var ErrInvalidScriptFile = errors.New("could not read script file")
var ErrScriptNotRunnable = errors.New("could not run script")

func RunPrompt() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for s.Scan() {
		line := s.Text()
		Run(line)
		// TODO: Understand the implications of this. The book implies that it's
		// to allow the users to keep issuing commands even if they make a mistake.
		lerrors.HadError = false
		fmt.Print("> ")
	}
}

func RunFile(path string) error {
	fBytes, err := os.ReadFile(path)
	if err != nil {
		return errors.Join(ErrInvalidScriptFile, err)
	}
	Run(string(fBytes))
	if lerrors.HadError {
		return ErrScriptNotRunnable
	}
	return nil
}

func Run(source string) {
	s := NewScanner(source)
	tokens := s.ScanTokens()

	for _, t := range tokens {
		fmt.Println(t)
	}
}
