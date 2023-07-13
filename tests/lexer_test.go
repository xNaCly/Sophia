package tests

import (
	"log"
	"os"
	"testing"
	"tisp/core"
)

func TestLexerHelloWorld(t *testing.T) {
	in := []byte(`[putv "Hello World!"]`)
	l := core.NewLexer(in)
	token := l.Lex()
	if len(token) == 0 {
		t.Error("Lexer found error, token empty")
	}

	expected := []int{
		core.LEFT_BRACE,
		core.PUTV,
		core.STRING,
		core.RIGHT_BRACE,
		core.EOF,
	}

	for i, tok := range token {
		if tok.Type != expected[i] {
			t.Errorf("given token '%+v' of type '%d' at pos '%d' does not match expected token '%d'\n", tok, tok.Type, i, expected[i])
		}
	}
}

func TestLexerFloats(t *testing.T) {
	null, _ := os.Open(os.DevNull)
	os.Stdout = null
	log.SetOutput(null)
	in := map[string]float64{
		"10.0":      10.0,
		"1_000_000": 1_000_000.0,
		".01":       0.01,
		".1e-3":     0.0001,
		"1.2e-2":    0.012,
		"15e4":      150_000,
	}
	for k, v := range in {
		l := core.NewLexer([]byte(k))
		o := l.Lex()
		if l.HasError {
			t.Fatalf("failed to lex float for input '%s', expected '%f'\n", k, v)
		}
		if o[0].Float != v {
			t.Fatalf("lexed float '%f' not equal to '%f' for input '%s'\n", o[0].Float, v, k)
		}
	}
}

func TestLexerArithmetic(t *testing.T) {
	in := []byte(`[add 1 [mul 1 [div 10 2]]]`)
	l := core.NewLexer(in)
	token := l.Lex()
	if len(token) == 0 {
		t.Error("Lexer found error, token empty")
	}

	expected := []int{
		core.LEFT_BRACE,
		core.ADD,
		core.FLOAT,
		core.LEFT_BRACE,
		core.MUL,
		core.FLOAT,
		core.LEFT_BRACE,
		core.DIV,
		core.FLOAT,
		core.FLOAT,
		core.RIGHT_BRACE,
		core.RIGHT_BRACE,
		core.RIGHT_BRACE,
		core.EOF,
	}

	for i, tok := range token {
		if tok.Type != expected[i] {
			t.Errorf("given token '%+v' of type '%d' at pos '%d' does not match expected token '%d'", tok, tok.Type, i, expected[i])
		}
	}
}

func TestLexerIgnoreCharsAndComments(t *testing.T) {
	in := []string{
		";;",
		";;comment",
		"\t\n ",
		"      ",
	}
	for _, v := range in {
		l := core.NewLexer([]byte(v))
		o := l.Lex()
		if len(o) != 1 {
			t.Errorf("lexer output for '%s' should be empty due to a comment, but contains '%v' of size '%d'", v, o, len(o))
		}
	}
}

func TestLexerErrorsOnUnknownTokenAndIntegers(t *testing.T) {
	null, _ := os.Open(os.DevNull)
	os.Stdout = null
	log.SetOutput(null)
	in := []string{
		"[t]",
		`;;comment
[putc "test?"]`,
	}
	for _, v := range in {
		l := core.NewLexer([]byte(v))
		o := l.Lex()
		if len(o) != 1 {
			t.Errorf("lexer output for '%s' should be empty due to a syntax error, but contains '%v' of size '%d'", v, o, len(o))
		}
	}
}
