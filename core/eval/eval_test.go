package eval

import (
	"sophia/core"
	"sophia/core/lexer"
	"sophia/core/parser"
	"sophia/core/serror"
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
			serror.SetDefault(serror.NewFormatter(&core.CONF, i.str, "test"))
			l := lexer.New(i.str)
			p := parser.New(l.Lex(), "test")
			r := Eval("repl", p.Parse())

			if serror.HasErrors() {
				t.Errorf("lexer or parser error for %q", i.str)
			}
			if len(r) == 0 {
				t.Errorf("eval result empty for %q", i.str)
				return
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
			str: "(let a 6)",
			exp: "6",
		},
		{
			str: "(let b 1 2 3)",
			exp: "[1 2 3]",
		},
		{
			str: "(let c (* 5 5))",
			exp: "25",
		},
		{
			str: "(let d (let e (+ 5 5)))",
			exp: "10",
		},
		{
			str: "(let f true)",
			exp: "true",
		},
		{
			str: "(let g false)",
			exp: "false",
		},
	}
	for _, i := range input {
		t.Run(i.str, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, i.str, "test"))
			l := lexer.New(i.str)
			p := parser.New(l.Lex(), "test")
			r := Eval("repl", p.Parse())
			if serror.HasErrors() {
				t.Errorf("lexer or parser error for %q", i.str)
			}
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
			str: "(if true)",
			exp: "true",
		},
		{
			str: "(if true (let a 5))",
			exp: "true",
		},
		{
			str: "(eq 1 2)",
			exp: "false",
		},
		{
			str: "(eq true false)",
			exp: "false",
		},
		{
			str: "(lt 10 1)",
			exp: "false",
		},
		{
			str: "(gt 1 10)",
			exp: "false",
		},
		{
			str: "(lt 1 10)",
			exp: "true",
		},
		{
			str: "(gt 10 1)",
			exp: "true",
		},
		{
			str: "(and true true)",
			exp: "true",
		},
		{
			str: "(and false true)",
			exp: "false",
		},
		{
			str: "(and false false)",
			exp: "false",
		},
		{
			str: "(or true true)",
			exp: "true",
		},
		{
			str: "(or false true)",
			exp: "true",
		},
		{
			str: "(or false false)",
			exp: "false",
		},
		{
			str: "(not false)",
			exp: "true",
		},
		{
			str: "(not true)",
			exp: "false",
		},
	}
	for _, i := range input {
		t.Run(i.str, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, i.str, "test"))
			l := lexer.New(i.str)
			p := parser.New(l.Lex(), "test")
			r := Eval("repl", p.Parse())
			if serror.HasErrors() {
				t.Errorf("lexer or parser error for %q", i.str)
			}
			if len(r) == 0 {
				t.Errorf("eval result empty")
			}
			if i.exp != r[0] {
				t.Errorf("%q not equal to %q", i.exp, r[0])
			}
		})
	}
}

func TestEvalMerge(t *testing.T) {
	input := []struct {
		str string
		exp string
	}{
		{
			str: `(++ "hello" "world")`,
			exp: "helloworld",
		},
		{
			str: `(let a 1 2)(++ a 1 2)`,
			exp: "[1 2 1 2]",
		},
		{
			str: `(++ "hello" 1 2)`,
			exp: "[hello 1 2]",
		},
		{
			str: `(++ 1 2 "hello")`,
			exp: "[1 2 hello]",
		},
	}
	for _, i := range input {
		t.Run(i.str, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, i.str, "test"))
			l := lexer.New(i.str)
			p := parser.New(l.Lex(), "test")
			r := Eval("repl", p.Parse())
			if serror.HasErrors() {
				t.Errorf("lexer or parser error for %q", i.str)
			}
			if len(r) == 0 {
				t.Errorf("eval result empty for %q", i.str)
				return
			}
			got := r[len(r)-1]
			if i.exp != got {
				t.Errorf("got %q, wanted %q", got, i.exp)
			}
		})
	}
}

func TestEvalFunction(t *testing.T) {
	input := []struct {
		str string
		exp string
	}{
		{
			str: "(fun square (_ a) (* a a))(square 12)",
			exp: "144",
		},
		{
			str: "(fun sum (_ a b) (+ a b))(sum 12 12)",
			exp: "24",
		},
		{
			str: "(fun print (_ a) (put a))(let y 12 23 12)(print y)",
			exp: "<nil>",
		},
	}
	for _, i := range input {
		t.Run(i.str, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, i.str, "test"))
			l := lexer.New(i.str)
			p := parser.New(l.Lex(), "test")
			r := Eval("repl", p.Parse())
			if serror.HasErrors() {
				t.Errorf("lexer or parser error for %q", i.str)
			}
			if len(r) == 0 {
				t.Errorf("eval result empty for %q", i.str)
				return
			}
			got := r[len(r)-1]
			if i.exp != got {
				t.Errorf("got %q, wanted %q", got, i.exp)
			}
		})
	}
}

func TestEvalLoop(t *testing.T) {
	input := []struct {
		str string
		exp string
	}{
		{
			str: "(let sum 0)(let arr 9)(for (_ e) arr (let sum (+ e sum)))(let r sum)",
			exp: "36",
		},
	}
	for _, i := range input {
		t.Run(i.str, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, i.str, "test"))
			l := lexer.New(i.str)
			p := parser.New(l.Lex(), "test")
			r := Eval("repl", p.Parse())
			if serror.HasErrors() {
				t.Errorf("lexer or parser error for %q", i.str)
			}
			if len(r) == 0 {
				t.Errorf("eval result empty for %q", i.str)
				return
			}
			got := r[len(r)-1]
			if i.exp != got {
				t.Errorf("got %q, wanted %q", got, i.exp)
			}
		})
	}
}

func TestEvalReturn(t *testing.T) {
	input := []struct {
		str string
		exp string
	}{
		{
			str: "(fun square (_n) (*n n))(square 12)",
			exp: "144",
		},
		{
			str: "(fun square (_n) (return (*n n)))(square 12)",
			exp: "144",
		},
		{
			str: `
(fun test (_ n) 
    (if (lt n 0) 
        (return -1))
    (return (*n n)
))
(test -12)
`,
			exp: "-1",
		},
	}
	for _, i := range input {
		t.Run(i.str, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, i.str, "test"))
			l := lexer.New(i.str)
			p := parser.New(l.Lex(), "test")
			r := Eval("repl", p.Parse())
			if serror.HasErrors() {
				t.Errorf("lexer or parser error for %q", i.str)
			}
			if len(r) == 0 {
				t.Errorf("eval result empty for %q", i.str)
				return
			}
			got := r[len(r)-1]
			if i.exp != got {
				t.Errorf("got %q, wanted %q", got, i.exp)
			}
		})
	}
}

func TestEvalObject(t *testing.T) {
	input := []struct {
		str string
		exp string
	}{
		{
			str: `
(let person {
    name: "anon"
    bank: {
        money: 2500
        institute: {
            name: "western union"
        }
    }
    age: 25
})
(let bankName person["bank"]["institute"]["name"])
            `,
			exp: "western union",
		},
		{
			str: `
(let person {
    name: "anon"
    bank: {
        money: 2500
        institute: {
            name: "western union"
        }
    }
    age: 25
})
(let arr person["bank"] 2 3 4)
(let money arr[0]["money"])
		`,
			exp: "2500",
		},
	}
	for _, i := range input {
		t.Run(i.str, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, i.str, "test"))
			l := lexer.New(i.str)
			p := parser.New(l.Lex(), "test")
			r := Eval("repl", p.Parse())
			if serror.HasErrors() {
				t.Errorf("lexer or parser error for %q", i.str)
			}
			if len(r) == 0 {
				t.Errorf("eval result empty for %q", i.str)
				return
			}
			got := r[len(r)-1]
			if i.exp != got {
				t.Errorf("got %q, wanted %q", got, i.exp)
			}
		})
	}
}

func TestEvalArray(t *testing.T) {
	input := []struct {
		str string
		exp string
	}{
		{
			str: `(let b #[1 2 3 4 5])(let r b[1])`,
			exp: "2",
		},
		{
			str: `(let b 5 #[1 2 3 4 5])(let r b[1][0])`,
			exp: "1",
		},
	}
	for _, i := range input {
		t.Run(i.str, func(t *testing.T) {
			serror.SetDefault(serror.NewFormatter(&core.CONF, i.str, "test"))
			l := lexer.New(i.str)
			p := parser.New(l.Lex(), "test")
			r := Eval("repl", p.Parse())
			if serror.HasErrors() {
				t.Errorf("lexer or parser error for %q", i.str)
			}
			if len(r) == 0 {
				t.Errorf("eval result empty for %q", i.str)
				return
			}
			got := r[len(r)-1]
			if i.exp != got {
				t.Errorf("got %q, wanted %q", got, i.exp)
			}
		})
	}
}
