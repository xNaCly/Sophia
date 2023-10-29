// Provides functionality for bubbling internal errors up to the user in a
// pretty and concise way
package serror

import (
	"fmt"
	"log"
	"sophia/core"
	"sophia/core/token"
	"strings"
)

const MAX_ERRORS = 5

type ErrorFormatter struct {
	Conf    *core.Config
	Lines   []string
	Errors  []Error
	Builder *strings.Builder
}

func (e *ErrorFormatter) HasErrors() bool {
	return len(e.Errors) > 0
}

func (e *ErrorFormatter) Add(t *token.Token, title string, info string, additional ...any) {
	e.Errors = append(e.Errors, Error{t, title, fmt.Sprintf(info, additional...)})
}

func (e *ErrorFormatter) Display() {
	for i, err := range e.Errors {
		if i < MAX_ERRORS && !e.Conf.AllErrors {
			log.Println(err.prettyPrint(e.Builder))
		} else {
			log.Printf("More than %d errors, skipping the remaining %d, rerun with '-all-errors' to view all", MAX_ERRORS, len(e.Errors)-MAX_ERRORS)
			break
		}
	}
}

type Error struct {
	Token *token.Token
	Title string // smth like Unknown token
	Info  string // in depth information: expected 'x' got 'y'
}

func (e *Error) prettyPrint(builder *strings.Builder) string {
	str := builder.String()
	builder.Reset()
	return str
}

// writes rune 'r' 'n' times into 'builder'
func runeRepeat(builder *strings.Builder, r rune, n int) {
	if n <= 0 {
		return
	}
	for i := 0; i < n; i++ {
		builder.WriteRune(r)
	}
}
