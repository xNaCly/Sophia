package run

import (
	"encoding/json"
	"errors"
	"fmt"
	"sophia/core"
	"sophia/core/debug"
	"sophia/core/eval"
	"sophia/core/lexer"
	"sophia/core/parser"
	"sophia/core/serror"
	"strings"
)

func run(input []byte, filename string) (s []string, e error) {
	defer func() {
		if core.CONF.Debug {
			return
		}
		if err := recover(); err != nil {
			return
		}
	}()
	errorFmt := serror.ErrorFormatter{
		Conf:    &core.CONF,
		Lines:   strings.Split(string(input), "\n"),
		Errors:  make([]serror.Error, 0),
		Builder: &strings.Builder{},
	}
	debug.Log("starting lexer")
	l := lexer.New(input, &errorFmt)
	tokens := l.Lex()
	debug.Log("lexed", len(tokens), "token")

	if core.CONF.Debug {
		debug.Log(debug.Token(tokens))
	}

	debug.Log("starting parser")
	p := parser.New(tokens, filename, &errorFmt)
	ast := p.Parse()

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
		debug.Log("done parsing - starting eval")
		s = eval.Eval(ast)
	}
	return
}
