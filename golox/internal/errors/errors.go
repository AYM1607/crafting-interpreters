package errors

import "fmt"

var HadError = false

func EmitError(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Printf(
		"[%d] Error%s: %s\n",
		line,
		where,
		message,
	)
	HadError = true
}
