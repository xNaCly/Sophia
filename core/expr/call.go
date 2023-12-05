package expr

import (
	"sophia/core/consts"
	"sophia/core/serror"
	"sophia/core/token"
	"strings"
)

type Call struct {
	Token  *token.Token
	Key    uint32
	Params []Node
}

func (c *Call) GetChildren() []Node {
	return c.Params
}

func (n *Call) SetChildren(c []Node) {
	n.Params = c
}

func (c *Call) GetToken() *token.Token {
	return c.Token
}

func (c *Call) Eval() any {
	storedFunc, ok := consts.FUNC_TABLE[c.Key]
	if !ok {
		serror.Add(c.Token, "Undefined function", "Function %q not defined", c.Token.Raw)
		serror.Panic()
	}

	def := castPanicIfNotType[*Func](storedFunc, c.Token)
	defParams := castPanicIfNotType[*Params](def.Params, c.Token)
	children := defParams.Children

	lchild := len(children)
	if len(children) != len(c.Params) {
		argLen := len(c.Params)
		if lchild < argLen {
			serror.Add(c.Token, "Too many arguments", "Too many arguments for %q, wanted %d, got %d", c.Token.Raw, lchild, len(c.Params))
			serror.Panic()
		} else if lchild > argLen {
			serror.Add(c.Token, "Not enough arguments", "Not enough arguments for %q, wanted %d, got %d", c.Token.Raw, lchild, len(c.Params))
			serror.Panic()
		}
	}

	// store variable values from before entering the function scope
	for i, arg := range c.Params {
		name := children[i].(*Ident)
		if val, ok := consts.SYMBOL_TABLE[name.Key]; ok {
			consts.SCOPE_TABLE[name.Key] = val
		}
		consts.SYMBOL_TABLE[name.Key] = arg.Eval()
	}

	var ret any

	for i, stmt := range def.Body {
		// enabling early returns
		if consts.RETURN.HasValue {
			ret = consts.RETURN.Value
			consts.RETURN.HasValue = false
			consts.RETURN.Value = nil
			break
		}
		if i+1 == len(def.Body) {
			ret = stmt.Eval()
			break
		}
		stmt.Eval()
	}

	// if last line was a return
	if consts.RETURN.HasValue {
		ret = consts.RETURN.Value
		consts.RETURN.HasValue = false
		consts.RETURN.Value = nil
	}

	defer func() {
		// going out of scope, therefore we restore variables used in the
		// function scope to their previous value stored in the local scope table
		for k, v := range consts.SCOPE_TABLE {
			consts.SYMBOL_TABLE[k] = v
			delete(consts.SCOPE_TABLE, k)
		}
	}()

	return ret
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
