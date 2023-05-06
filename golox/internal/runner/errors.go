package runner

import "fmt"

func emitError(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Printf(
		"[%d] Error%s: %s\n",
		line,
		where,
		message,
	)
	// TODO: The book sets `hadError` as true here, need to figure out where
	// that's used.
}
