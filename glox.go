package glox

type ErrorReporter interface {
	Process(line int, message string)
}
