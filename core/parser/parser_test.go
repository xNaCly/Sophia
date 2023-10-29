package parser

import (
	"log"
	"os"
	"sophia/core"
	"sophia/core/lexer"
	"sophia/core/serror"
	"strings"
	"testing"
)

func TestParserHelloWorld(t *testing.T) {
	in := []byte(`(put "Hello World!")`)
	errorFmt := serror.ErrorFormatter{
		Conf:    &core.CONF,
		Lines:   strings.Split(string(in), "\n"),
		Errors:  make([]serror.Error, 0),
		Builder: &strings.Builder{},
	}
	l := lexer.New([]byte(in), &errorFmt)
	token := l.Lex()

	New(token, "test", &errorFmt)
	if errorFmt.HasErrors() {
		t.Error("error while parsing hello world")
	}
}

func TestParserErrors(t *testing.T) {
	null, _ := os.Open(os.DevNull)
	os.Stdout = null
	log.SetOutput(null)
	in := []string{
		". darvin)",
		"(. 1",
		"(. 1 +)",
		"()",
		"((",
		"(()()())",
		"(/ () 12)",
		"(+ -)",
		"(+ (())",
	}
	for _, s := range in {
		t.Run(s, func(t *testing.T) {
			errorFmt := serror.ErrorFormatter{
				Conf:    &core.CONF,
				Lines:   strings.Split(string(s), "\n"),
				Errors:  make([]serror.Error, 0),
				Builder: &strings.Builder{},
			}
			l := lexer.New([]byte(s), &errorFmt)
			p := New(l.Lex(), "test", &errorFmt)
			p.Parse()
			if errorFmt.HasErrors() {
				t.Errorf("parsing should fail for %q, it did not", s)
			}
		})
	}
}
