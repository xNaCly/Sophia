package eval

import (
	"sophia/core/lexer"
	"sophia/core/parser"
	"testing"
)

func TestEvalAritmetic(t *testing.T) {
	input := []struct {
		str string
		exp string
	}{
		{
			str: "(+ 6 6)",
			exp: "12",
		},
		{
			str: "(+ (- 3 3) 6)",
			exp: "6",
		},
		{
			str: "(- (* 3 3) 6)",
			exp: "3",
		},
		{
			str: "(/ 0 (- 5 5))",
			exp: "NaN",
		},
		{
			str: "(/ 1 0)",
			exp: "+Inf",
		},
		{
			str: "(* (+ 6 (- 13 3)) (* 2 (+ 2 (- 8 2))))",
			exp: "256",
		},
	}
	for _, i := range input {
		t.Run(i.str, func(t *testing.T) {
			l := lexer.New([]byte(i.str))
			p := parser.NewParser(l.Lex())
			r := Eval(p.Parse())
			if len(r) == 0 {
				t.Errorf("eval result empty")
			}
			if i.exp != r[0] {
				t.Errorf("%q not equal to %q", i.exp, r[0])
			}
		})
	}
}

func TestEvalVariables(t *testing.T) {
	input := []struct {
		str string
		exp string
	}{
		{
			str: "(: a 6)",
			exp: "[6]",
		},
		{
			str: "(: b 1 2 3)",
			exp: "[1 2 3]",
		},
		{
			str: "(: c (* 5 5))",
			exp: "[25]",
		},
		{
			str: "(: d (: e (+ 5 5)))",
			exp: "[[10]]",
		},
		{
			str: "(: f true)",
			exp: "[true]",
		},
		{
			str: "(: g false)",
			exp: "[false]",
		},
	}
	for _, i := range input {
		t.Run(i.str, func(t *testing.T) {
			l := lexer.New([]byte(i.str))
			p := parser.NewParser(l.Lex())
			r := Eval(p.Parse())
			if len(r) == 0 {
				t.Errorf("eval result empty")
			}
			if i.exp != r[0] {
				t.Errorf("%q not equal to %q", i.exp, r[0])
			}
		})
	}
}

func TestEvalConditional(t *testing.T) {
	input := []struct {
		str string
		exp string
	}{
		{
			str: "(? true (. 1))",
			exp: "<nil>",
		},
		{
			str: "(? true (: a 5)(. a))",
			exp: "<nil>",
		},
		{
			str: "(= 1 2)",
			exp: "false",
		},
		{
			str: "(= true false)",
			exp: "false",
		},
		{
			str: "(& true true)",
			exp: "true",
		},
		{
			str: "(& false true)",
			exp: "false",
		},
		{
			str: "(& false false)",
			exp: "false",
		},
		{
			str: "(| true true)",
			exp: "true",
		},
		{
			str: "(| false true)",
			exp: "true",
		},
		{
			str: "(| false false)",
			exp: "false",
		},
		{
			str: "(! false)",
			exp: "true",
		},
		{
			str: "(! true)",
			exp: "false",
		},
	}
	for _, i := range input {
		t.Run(i.str, func(t *testing.T) {
			l := lexer.New([]byte(i.str))
			p := parser.NewParser(l.Lex())
			r := Eval(p.Parse())
			if len(r) == 0 {
				t.Errorf("eval result empty")
			}
			if i.exp != r[0] {
				t.Errorf("%q not equal to %q", i.exp, r[0])
			}
		})
	}
}
