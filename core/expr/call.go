package expr

import (
	"sophia/core/consts"
	"sophia/core/serror"
	"sophia/core/token"
	"strings"
)

type Call struct {
	Token  token.Token
	Params []Node
}

func (c *Call) GetToken() token.Token {
	return c.Token
}

func (c *Call) Eval() any {
	oldSymbols := map[string]any{}
	for k, v := range consts.SYMBOL_TABLE {
		oldSymbols[k] = v
	}

	storedFunc, ok := consts.FUNC_TABLE[c.Token.Raw]
	if !ok {
		serror.Add(&c.Token, "Undefined function", "Function %q not defined", c.Token.Raw)
		serror.Panic()
	}

	def := castPanicIfNotType[*Func](storedFunc, c.Token)
	defParams := castPanicIfNotType[*Params](def.Params, c.Token)

	if len(defParams.Children) != len(c.Params) {
		paramLen := len(defParams.Children)
		argLen := len(c.Params)
		if paramLen < argLen {
			serror.Add(&c.Token, "Too many arguments", "Too many arguments for %q, wanted %d, got %d", c.Token.Raw, len(defParams.Children), len(c.Params))
			serror.Panic()
		} else if paramLen > argLen {
			serror.Add(&c.Token, "Not enough arguments", "Not enough arguments for %q, wanted %d, got %d", c.Token.Raw, len(defParams.Children), len(c.Params))
			serror.Panic()
		}
	}

	for i, arg := range c.Params {
		consts.SYMBOL_TABLE[defParams.Children[i].GetToken().Raw] = arg.Eval()
	}

	// INFO: going out of scope, therefore we restore the previous state of the
	// symbol table, due to the fact that we disallow functions with side
	// effects
	// TODO: maybe implement this similary to the for loop implementation?
	defer func() {
		consts.SYMBOL_TABLE = oldSymbols
	}()

	for i, stmt := range def.Body {
		if i+1 == len(def.Body) {
			return stmt.Eval()
		}
		stmt.Eval()
	}

	return nil
}

func (n *Call) CompileJs(b *strings.Builder) {
	cLen := len(n.Params)
	b.WriteString(n.Token.Raw)
	b.WriteRune('(')
	for i, c := range n.Params {
		c.CompileJs(b)
		if i+1 < cLen {
			b.WriteRune(',')
		}
	}
	b.WriteRune(')')
}
