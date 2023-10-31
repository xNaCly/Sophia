// Provides functionality for bubbling internal errors up to the user in a
// pretty and concise way
package serror

import (
	"fmt"
	"log"
	"path/filepath"
	"sophia/core"
	"sophia/core/token"
	"strconv"
	"strings"
)

const MAX_ERRORS = 3

const (
	ANSI_RESET   = "\033[0m"
	ANSI_RED     = "\033[91m"
	ANSI_YELLOW  = "\033[93m"
	ANSI_BLUE    = "\033[94m"
	ANSI_MAGENTA = "\033[95m"
)

type ErrorFormatter struct {
	conf    *core.Config
	lines   []string
	errors  []Error
	builder *strings.Builder
	file    string
}

func (e *ErrorFormatter) HasErrors() bool {
	return len(e.errors) > 0
}

func (e *ErrorFormatter) Add(t *token.Token, title string, info string, additional ...any) {
	e.errors = append(e.errors, Error{t, title, fmt.Sprintf(info, additional...)})
}

func (e *ErrorFormatter) Display() {
	// if a file, lookup absolute path
	if e.file != "cli" && e.file != "repl" {
		// only errors on os.Getwd (we don't really care)
		path, _ := filepath.Abs(e.file)
		e.file = path
	}
	for i, err := range e.errors {
		if i < MAX_ERRORS || e.conf.AllErrors {
			log.Println(err.prettyPrint(e, e.builder))
			if i+1 != len(e.errors) {
				log.Println()
			}
		} else {
			log.Printf("Too many errors, skipping the remaining %d, rerun with '-all-errors' to view all %d errors", len(e.errors)-MAX_ERRORS, len(e.errors))
			break
		}
	}
}

type Error struct {
	Token *token.Token
	Title string // smth like Unknown token
	Info  string // in depth information: expected 'x' got 'y'
}

// responsible for formatting the error title and the  filename + line + pos
func (e *Error) title(errorFormatter *ErrorFormatter, builder *strings.Builder) {
	builder.WriteString(ANSI_RED)
	builder.WriteString("error: ")
	builder.WriteString(ANSI_RESET)
	builder.WriteString(e.Title)
	builder.WriteString("\n\n\tat: ")
	builder.WriteString(ANSI_BLUE)
	builder.WriteString(errorFormatter.file)
	builder.WriteString(ANSI_RESET)
	builder.WriteRune(':')
	builder.WriteString(strconv.Itoa(e.Token.Line + 1))
	builder.WriteRune(':')
	builder.WriteString(strconv.Itoa(e.Token.LinePos + 1))
	builder.WriteRune(':')
}

// responsible for formatting the code snippet
func (e *Error) snippet(errorFormatter *ErrorFormatter, builder *strings.Builder) {
	t := e.Token
	prevLineAmount := 2
	nextLineAmount := 2
	if len(errorFormatter.lines) == 1 {
		prevLineAmount = 0
		nextLineAmount = 0
	}

	if t.Line == 0 {
		prevLineAmount = 0
	}
	lineIndex := t.Line - prevLineAmount
	if lineIndex < 0 {
		lineIndex = 0
	}

	prevLines := errorFormatter.lines[lineIndex:t.Line]

	for _, line := range prevLines {
		e.line(line, lineIndex, builder)
		lineIndex++
	}

	// print the offending line
	e.error(errorFormatter.lines[t.Line], builder)
	lineIndex++

	nextLineAmount = t.Line + 1 + nextLineAmount
	if nextLineAmount >= len(errorFormatter.lines)-1 {
		nextLineAmount = len(errorFormatter.lines) - 1
	}
	baseLine := t.Line + 1
	if baseLine >= len(errorFormatter.lines)-1 {
		baseLine = len(errorFormatter.lines) - 1
	}

	nextLines := errorFormatter.lines[baseLine:nextLineAmount]

	for _, line := range nextLines {
		e.line(line, lineIndex, builder)
		lineIndex++
	}
}

// formats the error line
func (e *Error) error(line string, builder *strings.Builder) {
	e.line(line, e.Token.Line, builder)
	builder.WriteString("\n\t")
	builder.WriteString(fmt.Sprintf("%5s| ", ""))
	runeRepeat(builder, ' ', e.Token.LinePos)
	builder.WriteString(ANSI_RED)
	runeRepeat(builder, '^', len(e.Token.Raw))
	builder.WriteString(ANSI_RESET)
}

// formats a singular line
func (e *Error) line(line string, lineNum int, builder *strings.Builder) {
	builder.WriteString("\n\t")
	builder.WriteString(fmt.Sprintf("%5d| ", lineNum+1))
	builder.WriteString(line)
}

func (e *Error) prettyPrint(errorFormatter *ErrorFormatter, builder *strings.Builder) string {
	e.title(errorFormatter, builder)
	builder.WriteString("\n")
	e.snippet(errorFormatter, builder)
	builder.WriteString("\n\n")
	builder.WriteString(e.Info)
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
