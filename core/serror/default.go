package serror

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/xnacly/sophia/core"
	"github.com/xnacly/sophia/core/token"
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
	err := defaultFormatter.errors[len(defaultFormatter.errors)-1]
	panic("sophia: " + err.Title + ": " + err.Info)
}

func NewFormatter(config *core.Config, input string, filename string, w io.Writer) *ErrorFormatter {
	if w == nil {
		w = os.Stdout
	}
	return &ErrorFormatter{
		conf:   config,
		lines:  strings.Split(input, "\n"),
		file:   filename,
		w:      bufio.NewWriter(w),
		errors: make([]Error, 0),
	}
}
