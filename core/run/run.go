package run

import (
	"encoding/json"
	"errors"
	"log"
	"sophia/core"
	"sophia/core/eval"
	"sophia/core/lexer"
	"sophia/core/parser"
)

func run(input []byte) (s []string, e error) {
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
	core.DbgLog("starting lexer")
	l := lexer.New(input)
	tokens := l.Lex()
	core.DbgLog("lexed", len(tokens), "token")
	if core.CONF.Debug {
		v, _ := json.MarshalIndent(tokens, "", "  ")
		core.DbgLog(string(v))
	}

	core.DbgLog("starting parser")
	p := parser.New(tokens)
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
		v, _ := json.MarshalIndent(ast, "", "  ")
		core.DbgLog(string(v))
	}
	core.DbgLog("done parsing - no errors, starting eval")
	s = eval.Eval(ast)
	return
}
