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
	reader        io.Reader
	writer        io.Writer
	errorReporter glox.ErrorReporter
}

func NewInterpreter(r io.Reader, w io.Writer, er glox.ErrorReporter) *Interpreter {
	return &Interpreter{
		reader:        r,
		writer:        w,
		errorReporter: er,
	}
}

func (i *Interpreter) Start(args []string) {
	if len(args) > 1 {
		fmt.Fprintln(i.writer, "Usage: glox [script]")
		os.Exit(64)
	}

	if len(args) == 1 {
		i.runFile(args[0])
		if i.errorReporter.HadError() {
			os.Exit(65)
		}
	} else {
		i.runPrompt()
	}
}

func (i *Interpreter) runFile(path string) error {
	source, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return i.run(string(source))
}

func (i *Interpreter) runPrompt() error {
	reader := bufio.NewReader(i.reader)
	for {
		fmt.Fprint(i.writer, "> ")
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
		i.errorReporter.Reset()
	}
	return nil
}

func (i *Interpreter) run(source string) error {
	scanner := NewScanner(source, i.errorReporter)
	tokens := scanner.ScanTokens()
	for _, token := range tokens {
		fmt.Fprintln(i.writer, token)
	}
	return nil
}
