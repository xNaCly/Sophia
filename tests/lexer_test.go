package tests

import (
	"log"
	"os"
	"sophia/core"
	"testing"
)

func TestLexerHelloWorld(t *testing.T) {
	in := []byte(`(. "Hello World!")`)
	l := core.NewLexer(in)
	token := l.Lex()
	if len(token) == 0 {
		t.Error("Lexer found error, token empty")
	}

	expected := []int{
		core.LEFT_BRACE,
		core.PUT,
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
	in := []string{
		"10.0",
		"1_000_000",
		"0.01",
		"0.1e-3",
		"1.2e-2",
		"15e4",
	}
	for _, v := range in {
		t.Run(v, func(t *testing.T) {
			l := core.NewLexer([]byte(v))
			o := l.Lex()
			if l.HasError || len(o) == 0 {
				t.Fatalf("failed to lex float for input '%s'\n", v)
			}
			if o[0].Type != core.FLOAT {
				t.Fatalf("'%s' was not lexed as a float, got %s", v, core.TOKEN_NAME_MAP[o[0].Type])
			}
		})
	}
}

func TestLexerIdent(t *testing.T) {
	in := []byte(`b a abc abcdefghijklmnopqrstuvwxyz`)
	l := core.NewLexer(in)
	token := l.Lex()
	if len(token) == 0 {
		t.Error("Lexer found error, token empty")
	}

	expectedType := []int{
		core.IDENT,
		core.IDENT,
		core.IDENT,
		core.IDENT,
		core.EOF,
	}

	expectedRaw := []string{
		"b",
		"a",
		"abc",
		"abcdefghijklmnopqrstuvwxyz",
		"",
	}

	for i, tok := range token {
		if tok.Type != expectedType[i] {
			t.Errorf("given token '%+v' of type '%d' at pos '%d' does not match expected token '%d'", tok, tok.Type, i, expectedType[i])
		}
		if tok.Raw != expectedRaw[i] {
			t.Errorf("given raw content '%s' at pos '%d' does not match expected content '%s'", tok.Raw, i, expectedRaw[i])
		}
	}
}

func TestLexerOperators(t *testing.T) {
	in := []byte(`.+-/*%:()?=|&`)
	l := core.NewLexer(in)
	token := l.Lex()
	if len(token) == 0 {
		t.Error("Lexer found error, token empty")
	}

	expected := []int{
		core.PUT,
		core.ADD,
		core.SUB,
		core.DIV,
		core.MUL,
		core.MOD,
		core.COLON,
		core.LEFT_BRACE,
		core.RIGHT_BRACE,
		core.IF,
		core.EQUAL,
		core.OR,
		core.AND,
		core.EOF,
	}

	for i, tok := range token {
		if tok.Type != expected[i] {
			t.Errorf("given token '%+v' of type '%d' at pos '%d' does not match expected token '%d'", tok, tok.Type, i, expected[i])
		}
	}
}

func TestLexerArithmetic(t *testing.T) {
	in := []byte(`(+ 1 (* 1 (/ 1 (% 1))))`)
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
		core.LEFT_BRACE,
		core.MOD,
		core.FLOAT,
		core.RIGHT_BRACE,
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
		t.Run(v, func(t *testing.T) {
			l := core.NewLexer([]byte(v))
			o := l.Lex()
			if len(o) != 1 {
				t.Errorf("lexer output for '%s' should be empty due to a comment, but contains '%v' of size '%d'", v, o, len(o))
			}
		})
	}
}

func TestLexerErrorsOnUnknownTokenAndIntegers(t *testing.T) {
	null, _ := os.Open(os.DevNull)
	os.Stdout = null
	log.SetOutput(null)
	in := []string{
		"!",
		`;;comment
?[putc "test"]`,
	}
	for _, v := range in {
		t.Run(v, func(t *testing.T) {
			l := core.NewLexer([]byte(v))
			o := l.Lex()
			if len(o) != 1 {
				t.Errorf("lexer output for '%s' should be empty due to a syntax error, but contains '%v' of size '%d'", v, o, len(o))
			}
		})
	}
}

func TestLexerBooleans(t *testing.T) {
	in := []byte(`true false`)
	l := core.NewLexer(in)
	token := l.Lex()
	if len(token) == 0 {
		t.Error("Lexer found error, token empty")
	}

	expectedType := []int{
		core.BOOL,
		core.BOOL,
		core.EOF,
	}

	expectedRaw := []string{
		"true",
		"false",
		"",
	}

	for i, tok := range token {
		if tok.Type != expectedType[i] {
			t.Errorf("given token '%+v' of type '%d' at pos '%d' does not match expected token '%d'", tok, tok.Type, i, expectedType[i])
		}
		if tok.Raw != expectedRaw[i] {
			t.Errorf("given raw content '%s' at pos '%d' does not match expected content '%s'", tok.Raw, i, expectedRaw[i])
		}
	}
}
