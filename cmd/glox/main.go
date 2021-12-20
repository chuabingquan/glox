package main

import (
	"glox/interpreter"
	"os"
)

func main() {
	errorReporter := interpreter.NewErrorReporter(os.Stdout)
	glox := interpreter.NewInterpreter(os.Stdin, os.Stdout, errorReporter)
	glox.Start(os.Args[1:])
}
