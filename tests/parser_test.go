package tests

import (
	"log"
	"os"
	"sophia/core"
	"testing"
)

func TestParserHelloWorld(t *testing.T) {
	in := []byte(`(. "Hello World!")`)
	l := core.NewLexer(in)
	token := l.Lex()

	p := core.NewParser(token)
	if l.HasError || p.HasError {
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
			l := core.NewLexer([]byte(s))
			p := core.NewParser(l.Lex())
			a := p.Parse()
			if !p.HasError || len(a) != 0 {
				t.Errorf("parsing should fail for %q, it did not", s)
			}
		})
	}
}
