package runner

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var ErrInvalidScriptFile = errors.New("could not read script file")

func RunPrompt() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for s.Scan() {
		line := s.Text()
		Run(line)
		// TODO: resed hadError wherever it is set.
		fmt.Print("> ")
	}
}

func RunFile(path string) error {
	fBytes, err := os.ReadFile(path)
	if err != nil {
		return errors.Join(ErrInvalidScriptFile, err)
	}
	Run(string(fBytes))
	// TODO: check hadError and exit with a 65 code if so.
	return nil
}

func Run(source string) {
	s := NewScanner(source)
	tokens := s.ScanTokens()

	for _, t := range tokens {
		fmt.Println(t)
	}
}
