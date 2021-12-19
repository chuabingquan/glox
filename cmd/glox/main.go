package main

import (
	"glox/interpreter"
	"os"
)

func main() {
	errorReporter := interpreter.NewErrorReporter()
	glox := interpreter.NewInterpreter(errorReporter)
	glox.Start(os.Args[1:])
}
