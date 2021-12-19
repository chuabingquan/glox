package interpreter

import "fmt"

type ErrorReporter struct{}

func NewErrorReporter() *ErrorReporter {
	return &ErrorReporter{}
}

func (er ErrorReporter) Process(line int, message string) {
	er.report(line, "", message)
}

func (er ErrorReporter) report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
}
