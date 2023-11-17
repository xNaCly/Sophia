package run

import (
	"encoding/json"
	"errors"
	"fmt"
	"sophia/core"
	"sophia/core/debug"
	"sophia/core/eval"
	"sophia/core/lexer"
	"sophia/core/optimizer"
	"sophia/core/parser"
	"sophia/core/serror"
)

func run(input string, filename string) (s []string, e error) {
	defer func() {
		if core.CONF.Debug {
			return
		}
		if err := recover(); err != nil {
			serror.Display()
			return
		}
	}()

	serror.SetDefault(serror.NewFormatter(&core.CONF, input, filename))

	debug.Log("starting lexer")
	l := lexer.New(input)
	tokens := l.Lex()
	if serror.HasErrors() {
		serror.Display()
		e = errors.New("Syntax errors found, skipping remaining interpreter stages. (parsing and evaluation)")
		return
	}
	debug.Log("lexed", len(tokens), "token")

	if core.CONF.Debug {
		debug.Log(debug.Token(tokens))
	}

	debug.Log("starting parser")
	p := parser.New(tokens, filename)
	ast := p.Parse()
	if serror.HasErrors() {
		serror.Display()
		e = errors.New("Semantic errors found, skipping remaining interpreter stages. (evaluation)")
		return
	}

	if core.CONF.Debug {
		out, _ := json.MarshalIndent(ast, "", "  ")
		debug.Log(string(out))
	}

	if filename != "repl" {
		if core.CONF.EnableOptimizer {
			debug.Log("done parsing - starting optimizer")
			opt := optimizer.New()
			ast = opt.Start(ast)
			if core.CONF.Debug {
				out, _ := json.MarshalIndent(ast, "", "  ")
				debug.Log(string(out))
			}
		}
	} else {
		debug.Log("done parsing - starting eval")
	}

	if len(ast) == 0 {
		return
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
		s = eval.Eval(filename, ast)
		debug.Log("done evaling")
	}

	return
}
