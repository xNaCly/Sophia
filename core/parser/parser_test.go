package parser

import (
	"strings"
	"testing"

	"github.com/xnacly/sophia/core"
	"github.com/xnacly/sophia/core/lexer"
	"github.com/xnacly/sophia/core/serror"
)

func TestParserHelloWorld(t *testing.T) {
	in := `(println "Hello World!")`

	serror.SetDefault(serror.NewFormatter(&core.CONF, in, "test", nil))
	l := lexer.New(strings.NewReader(in))
	token := l.Lex()

	New(token, "test")
	if serror.HasErrors() {
		t.Error("error while parsing hello world")
	}
}

func TestParserErrors(t *testing.T) {
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
			serror.SetDefault(serror.NewFormatter(&core.CONF, s, "test", nil))
			l := lexer.New(strings.NewReader(s))
			p := New(l.Lex(), "test")
			p.Parse()
			if !serror.HasErrors() {
				t.Errorf("parsing should fail for %q, it did not", s)
			}
		})
	}
}

func TestParserIndex(t *testing.T) {
	in := []string{
		`(println person#["name"])`,
		`(println person#["data"]["name"])`,
		`(println person#["data"]["name"][0])`,
	}
	for _, s := range in {
		t.Run(s, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, s, "test", nil))
			l := lexer.New(strings.NewReader(s))
			tokens := l.Lex()
			p := New(tokens, "test")
			p.Parse()
			if serror.HasErrors() {
				serror.Display()
				t.Errorf("parsing should not fail for %q, it did", s)
			}
		})
	}
}

func TestParserArray(t *testing.T) {
	in := []string{
		"(let array [1 2 3 4 5])",
		"(for [i] [1 2 3 4 5])",
		"(++ [1 2 3 4 5] 1 2)",
	}
	for _, s := range in {
		t.Run(s, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, s, "test", nil))
			l := lexer.New(strings.NewReader(s))
			tokens := l.Lex()
			p := New(tokens, "test")
			p.Parse()
			if serror.HasErrors() {
				serror.Display()
				t.Errorf("parsing should not fail for %q, it did", s)
			}
		})
	}
}

func TestParserModules(t *testing.T) {
	in := []string{
		"(module person)",
		"(module person (fun str [p] (++ \"person: \" p#[\"name\"])))",
	}

	for _, s := range in {
		t.Run(s, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, s, "test", nil))
			l := lexer.New(strings.NewReader(s))
			tokens := l.Lex()
			p := New(tokens, "test")
			p.Parse()
			if serror.HasErrors() {
				serror.Display()
				t.Errorf("parsing should not fail for %q, it did", s)
			}
		})
	}
}
