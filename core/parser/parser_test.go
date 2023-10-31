package parser

import (
	"log"
	"os"
	"sophia/core"
	"sophia/core/lexer"
	"sophia/core/serror"
	"testing"
)

func TestParserHelloWorld(t *testing.T) {
	in := `(put "Hello World!")`

	serror.SetDefault(serror.NewFormatter(&core.CONF, in, "test"))
	l := lexer.New(in)
	token := l.Lex()

	New(token, "test")
	if serror.HasErrors() {
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
			serror.SetDefault(serror.NewFormatter(&core.CONF, s, "test"))
			l := lexer.New(s)
			p := New(l.Lex(), "test")
			p.Parse()
			if !serror.HasErrors() {
				t.Errorf("parsing should fail for %q, it did not", s)
			}
		})
	}
}
