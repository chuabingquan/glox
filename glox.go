package glox

type ErrorReporter interface {
	HadError() bool
	Process(line int, message string)
	Reset()
}
