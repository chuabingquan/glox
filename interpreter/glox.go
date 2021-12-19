package interpreter

import (
	"bufio"
	"fmt"
	"glox"
	"io"
	"os"
)

type Interpreter struct {
	hadError      bool
	errorReporter glox.ErrorReporter
}

func NewInterpreter(er glox.ErrorReporter) *Interpreter {
	return &Interpreter{
		hadError:      false,
		errorReporter: er,
	}
}

func (i Interpreter) Start(args []string) {
	if len(args) > 1 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	}

	if len(args) == 1 {
		i.runFile(args[0])
		if i.hadError {
			os.Exit(65)
		}
	} else {
		i.runPrompt()
	}
}

func (i Interpreter) runFile(path string) error {
	source, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return i.run(string(source))
}

func (i Interpreter) runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		err = i.run(line)
		if err != nil {
			return err
		}
	}
	i.hadError = false
	return nil
}

func (i Interpreter) run(source string) error {
	return nil
}
