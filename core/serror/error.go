// Provides functionality for bubbling internal errors up to the user in a
// pretty and concise way
package serror

import (
	"bufio"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/xnacly/sophia/core"
	"github.com/xnacly/sophia/core/token"
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
	conf   *core.Config
	lines  []string
	errors []Error
	w      *bufio.Writer
	file   string
}

func (e *ErrorFormatter) HasErrors() bool {
	return len(e.errors) > 0
}

func (e *ErrorFormatter) Add(t *token.Token, title string, info string, additional ...any) {
	e.errors = append(e.errors, Error{t, title, fmt.Sprintf(info, additional...)})
}

func (e *ErrorFormatter) Display() {
	if len(e.errors) == 0 {
		return
	}
	// if a file, lookup absolute path
	if e.file != "cli" && e.file != "repl" && e.file != "stdin" {
		// only errors on os.Getwd (we don't really care)
		path, _ := filepath.Abs(e.file)
		e.file = path
	}
	for i, err := range e.errors {
		if i < MAX_ERRORS || e.conf.AllErrors {
			err.prettyPrint(e)
			if i+1 != len(e.errors) {
				e.w.WriteRune('\n')
			}
		} else {
			fmt.Fprintf(e.w, "Too many errors, skipping the remaining %d, rerun with '-all-errors' to view all %d errors", len(e.errors)-MAX_ERRORS, len(e.errors))
			break
		}
	}
	e.w.Flush()
}

type Error struct {
	Token *token.Token
	Title string // smth like Unknown token
	Info  string // in depth information: expected 'x' got 'y'
}

// responsible for formatting the error title and the  filename + line + pos
func (e *Error) title(errFmt *ErrorFormatter) {
	errFmt.w.WriteString(ANSI_RED)
	errFmt.w.WriteString("error: ")
	errFmt.w.WriteString(ANSI_RESET)
	errFmt.w.WriteString(e.Title)
	errFmt.w.WriteString("\n\n\tat: ")
	errFmt.w.WriteString(ANSI_BLUE)
	errFmt.w.WriteString(errFmt.file)
	errFmt.w.WriteString(ANSI_RESET)
	errFmt.w.WriteRune(':')
	errFmt.w.WriteString(strconv.Itoa(e.Token.Line + 1))
	errFmt.w.WriteRune(':')
	errFmt.w.WriteString(strconv.Itoa(e.Token.LinePos + 1))
	errFmt.w.WriteRune(':')
}

// responsible for formatting the code snippet
func (e *Error) snippet(errFmt *ErrorFormatter) {
	t := e.Token
	prevLineAmount := 2
	nextLineAmount := 2
	if len(errFmt.lines) == 1 {
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

	prevLines := errFmt.lines[lineIndex:t.Line]

	for _, line := range prevLines {
		e.line(errFmt, line, lineIndex)
		lineIndex++
	}

	// print the offending line
	e.error(errFmt, errFmt.lines[t.Line])
	lineIndex++

	nextLineAmount = t.Line + 1 + nextLineAmount
	if nextLineAmount >= len(errFmt.lines)-1 {
		nextLineAmount = len(errFmt.lines) - 1
	}
	baseLine := t.Line + 1
	if baseLine >= len(errFmt.lines)-1 {
		baseLine = len(errFmt.lines) - 1
	}

	nextLines := errFmt.lines[baseLine:nextLineAmount]

	for _, line := range nextLines {
		e.line(errFmt, line, lineIndex)
		lineIndex++
	}
}

// formats the error line
func (e *Error) error(errFmt *ErrorFormatter, line string) {
	e.line(errFmt, line, e.Token.Line)
	errFmt.w.WriteString("\n\t")
	fmt.Fprintf(errFmt.w, "%5s| ", " ")
	runeRepeat(errFmt.w, ' ', e.Token.LinePos-1)
	errFmt.w.WriteString(ANSI_RED)
	runeRepeat(errFmt.w, '^', len(e.Token.Raw))
	errFmt.w.WriteString(ANSI_RESET)
}

// formats a singular line
func (e *Error) line(errFmt *ErrorFormatter, line string, lineNum int) {
	errFmt.w.WriteString("\n\t")
	fmt.Fprintf(errFmt.w, "%5d| ", lineNum+1)
	errFmt.w.WriteString(line)
}

func (e *Error) prettyPrint(errFmt *ErrorFormatter) {
	e.title(errFmt)
	errFmt.w.WriteRune('\n')
	e.snippet(errFmt)
	errFmt.w.WriteRune('\n')
	errFmt.w.WriteRune('\n')
	errFmt.w.WriteString(e.Info)
	errFmt.w.WriteRune('\n')
}

// writes rune 'r' 'n' times into 'builder'
func runeRepeat(w *bufio.Writer, r rune, n int) {
	if n <= 0 {
		return
	}
	for i := 0; i < n; i++ {
		w.WriteRune(r)
	}
}
