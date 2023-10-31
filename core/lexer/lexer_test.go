package lexer

import (
	"fmt"
	"sophia/core"
	"sophia/core/serror"
	"sophia/core/token"
	"testing"
)

func TestLexerHelloWorld(t *testing.T) {
	in := `(put "Hello World!")`

	serror.SetDefault(serror.NewFormatter(&core.CONF, in, "test"))
	l := New(in)
	tok := l.Lex()
	if len(tok) == 0 {
		t.Error("Lexer found error, token empty")
	}

	expected := []int{
		token.LEFT_BRACE,
		token.PUT,
		token.STRING,
		token.RIGHT_BRACE,
		token.EOF,
	}

	for i, toke := range tok {
		if toke.Type != expected[i] {
			t.Errorf("given token '%+v' of type '%d' at pos '%d' does not match expected token '%d'\n", toke, toke.Type, i, expected[i])
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
			serror.SetDefault(serror.NewFormatter(&core.CONF, v, "test"))
			l := New(v)
			o := l.Lex()
			if serror.HasErrors() {
				t.Fatalf("failed to lex float for input '%s'\n", v)
			}
			if o[0].Type != token.FLOAT {
				t.Fatalf("'%s' was not lexed as a float, got %s", v, token.TOKEN_NAME_MAP[o[0].Type])
			}
		})
	}
}

func TestLexerIdent(t *testing.T) {
	in := `b a abc abcdefghijklmnopqrstuvwxyz`
	serror.SetDefault(serror.NewFormatter(&core.CONF, in, "test"))
	l := New(in)
	to := l.Lex()
	if serror.HasErrors() {
		t.Error("Lexer found error, token empty")
	}

	expectedType := []int{
		token.IDENT,
		token.IDENT,
		token.IDENT,
		token.IDENT,
		token.EOF,
	}

	expectedRaw := []string{
		"b",
		"a",
		"abc",
		"abcdefghijklmnopqrstuvwxyz",
		" ",
	}

	for i, tok := range to {
		if tok.Type != expectedType[i] {
			t.Errorf("given token '%+v' of type '%d' at pos '%d' does not match expected token '%d'", tok, tok.Type, i, expectedType[i])
		}
		if tok.Raw != expectedRaw[i] {
			t.Errorf("given raw content '%s' at pos '%d' does not match expected content '%s'", tok.Raw, i, expectedRaw[i])
		}
	}
}

func TestLexerOperators(t *testing.T) {
	in := `put +-/*% let () if eq or and not ++ fun _ for gt lt match`
	serror.SetDefault(serror.NewFormatter(&core.CONF, in, "test"))
	l := New(in)
	to := l.Lex()
	if serror.HasErrors() {
		t.Error("Lexer found error, token empty")
	}

	expected := []int{
		token.PUT,
		token.ADD,
		token.SUB,
		token.DIV,
		token.MUL,
		token.MOD,
		token.LET,
		token.LEFT_BRACE,
		token.RIGHT_BRACE,
		token.IF,
		token.EQUAL,
		token.OR,
		token.AND,
		token.NEG,
		token.MERGE,
		token.FUNC,
		token.PARAM,
		token.FOR,
		token.GT,
		token.LT,
		token.MATCH,
		token.EOF,
	}

	for i, toke := range to {
		if toke.Type != expected[i] {
			t.Errorf("given token '%+v' of type '%d' at pos '%d' does not match expected token '%d'", toke, toke.Type, i, expected[i])
		}
	}
}

func TestLexerArithmetic(t *testing.T) {
	in := `(+ 1 (* 1 (/ 1 (% 1))))`
	serror.SetDefault(serror.NewFormatter(&core.CONF, in, "test"))
	l := New(in)
	to := l.Lex()
	if serror.HasErrors() {
		t.Error("Lexer found error, token empty")
	}

	expected := []int{
		token.LEFT_BRACE,
		token.ADD,
		token.FLOAT,
		token.LEFT_BRACE,
		token.MUL,
		token.FLOAT,
		token.LEFT_BRACE,
		token.DIV,
		token.FLOAT,
		token.LEFT_BRACE,
		token.MOD,
		token.FLOAT,
		token.RIGHT_BRACE,
		token.RIGHT_BRACE,
		token.RIGHT_BRACE,
		token.RIGHT_BRACE,
		token.EOF,
	}

	for i, toke := range to {
		if toke.Type != expected[i] {
			t.Errorf("given token '%+v' of type '%d' at pos '%d' does not match expected token '%d'", toke, toke.Type, i, expected[i])
		}
	}
}

func TestLexerIgnoreCharsAndComments(t *testing.T) {
	in := []string{
		";;",
		";;comment",
		";;comment        \t\n",
	}
	for _, v := range in {
		t.Run(v, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, v, "test"))
			l := New(v)
			toks := []token.Token{}
			if l != nil {
				toks = l.Lex()
			}
			if serror.HasErrors() {
				t.Error("Lexer should have not found errors")
			}
			if len(toks) != 1 {
				fmt.Println(toks)
				t.Error("Lexer should have resulted in 1 token")
			}
		})
	}
}

func TestLexerErrorsOnUnknownTokenAndIntegers(t *testing.T) {
	in := []string{
		`;;comment
?[putc "test"]`,
	}
	for _, v := range in {
		t.Run(v, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, v, "test"))
			l := New(v)
			l.Lex()
			if !serror.HasErrors() {
				t.Error("Lexer should have found errors")
			}
		})
	}
}

func TestLexerBooleans(t *testing.T) {
	in := "true false"
	serror.SetDefault(serror.NewFormatter(&core.CONF, in, "test"))
	l := New(in)
	tok := l.Lex()
	if serror.HasErrors() {
		t.Error("Lexer found error, token empty")
	}

	expectedType := []int{
		token.BOOL,
		token.BOOL,
		token.EOF,
	}

	expectedRaw := []string{
		"true",
		"false",
		" ",
	}

	for i, toke := range tok {
		if toke.Type != expectedType[i] {
			t.Errorf("given token '%+v' of type '%d' at pos '%d' does not match expected token '%d'", toke, toke.Type, i, expectedType[i])
		}
		if toke.Raw != expectedRaw[i] {
			t.Errorf("given raw content '%s' at pos '%d' does not match expected content '%s'", toke.Raw, i, expectedRaw[i])
		}
	}
}
