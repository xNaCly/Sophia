package serror

import (
	"sophia/core"
	"sophia/core/token"
	"strings"
)

var defaultFormatter *ErrorFormatter

func Add(t *token.Token, title string, info string, additional ...any) {
	defaultFormatter.Add(t, title, info, additional...)
}

func Display() {
	defaultFormatter.Display()
}

func HasErrors() bool {
	return defaultFormatter.HasErrors()
}

func SetDefault(e *ErrorFormatter) {
	defaultFormatter = e
}

func Default() *ErrorFormatter {
	return defaultFormatter
}

func Panic() {
	panic("Runtime error")
}

func NewFormatter(config *core.Config, input string, filename string) *ErrorFormatter {
	return &ErrorFormatter{
		conf:    config,
		lines:   strings.Split(input, "\n"),
		file:    filename,
		builder: &strings.Builder{},
		errors:  make([]Error, 0),
	}
}
