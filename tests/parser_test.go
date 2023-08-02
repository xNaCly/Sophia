package tests

import (
	"encoding/json"
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

func TestParserArithmetic(t *testing.T) {
	in := []struct {
		inp string
		exp string
	}{
		{
			`(. "hello world")`,
			`[{"Token":{"Pos":1,"Line":0,"Type":8,"Raw":"","Float":0},"Children":[{"Token":{"Pos":5,"Line":0,"Type":3,"Raw":"hello world","Float":0}}]}]`,
		},
		{
			`(+ 1 2 (- 1 2))`,
			`[{"Token":{"Pos":1,"Line":0,"Type":4,"Raw":"","Float":0},"Children":[{"Token":{"Pos":3,"Line":0,"Type":2,"Raw":"1","Float":1}},{"Token":{"Pos":5,"Line":0,"Type":2,"Raw":"2","Float":2}},{"Token":{"Pos":8,"Line":0,"Type":5,"Raw":"","Float":0},"Children":[{"Token":{"Pos":10,"Line":0,"Type":2,"Raw":"1","Float":1}},{"Token":{"Pos":12,"Line":0,"Type":2,"Raw":"2","Float":2}}]}]}]`,
		},
	}
	for _, s := range in {
		t.Run(s.inp, func(t *testing.T) {
			l := core.NewLexer([]byte(s.inp))
			p := core.NewParser(l.Lex())
			v, _ := json.Marshal(p.Parse())
			if string(v) != string(s.exp) {
				t.Errorf("parsing %s not equal to expected result %s != %s", s.inp, string(v), s.exp)
			}
		})
	}
}
