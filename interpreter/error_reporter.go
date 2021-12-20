package interpreter

import (
	"fmt"
	"io"
)

type ErrorReporter struct {
	writer   io.Writer
	hadError bool
}

func NewErrorReporter(w io.Writer) *ErrorReporter {
	return &ErrorReporter{
		writer:   w,
		hadError: false,
	}
}

func (er *ErrorReporter) HadError() bool {
	return er.hadError
}

func (er *ErrorReporter) Process(line int, message string) {
	er.report(line, "", message)
}

func (er *ErrorReporter) Reset() {
	er.hadError = false
}

func (er *ErrorReporter) report(line int, where string, message string) {
	fmt.Fprintf(er.writer, "[line %d] Error %s: %s\n", line, where, message)
	er.hadError = true
}
