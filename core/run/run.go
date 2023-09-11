package run

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sophia/core"
	"sophia/core/debug"
	"sophia/core/eval"
	"sophia/core/lexer"
	"sophia/core/parser"
)

func run(input []byte, filename string) (s []string, e error) {
	defer func() {
		if core.CONF.Debug {
			return
		}
		if err := recover(); err != nil {
			log.Printf("err: %s", err)
			e = errors.New("runtime error")
			return
		}
	}()
	debug.Log("starting lexer")
	l := lexer.New(input)
	tokens := l.Lex()
	debug.Log("lexed", len(tokens), "token")

	if core.CONF.Debug {
		debug.Log(debug.Token(tokens))
	}

	debug.Log("starting parser")
	p := parser.New(tokens, filename)
	ast := p.Parse()
	if l.HasError {
		e = errors.New("lexer error")
		return
	}
	if p.HasError {
		e = errors.New("parser error")
		return
	}

	if core.CONF.Debug {
		out, _ := json.MarshalIndent(ast, "", "  ")
		debug.Log(string(out))
	}
	if len(core.CONF.Target) > 0 {
		trgt := core.CONF.Target
		debug.Log("done parsing - no errors, starting compilation for", trgt)
		if _, ok := core.TARGETS[trgt]; !ok {
			e = errors.New(fmt.Sprintf("compilation error: %q not found in compilation targets. Available targets: %s", trgt, core.TARGETS))
			return
		}
		fmt.Println(eval.CompileJs(ast))
	} else {
		debug.Log("done parsing - no errors, starting eval")
		s = eval.Eval(ast)
	}
	return
}
