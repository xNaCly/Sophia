package expr

import (
	"sophia/core/consts"
	"sophia/core/token"
"sophia/core/types"
	"strings"
)

// defining a variable
type Var struct {
	Token *token.Token
	Ident *Ident
	Value []types.Node
}

func (v *Var) GetChildren() []types.Node {
	return v.Value
}

func (n *Var) SetChildren(c []types.Node) {
	n.Value = c
}

func (v *Var) GetToken() *token.Token {
	return v.Token
}

func (v *Var) Eval() any {
	var val any
	if len(v.Value) > 1 {
		val = make([]any, len(v.Value))
		for i, c := range v.Value {
			val.([]any)[i] = c.Eval()
		}
	} else if len(v.Value) == 0 {
		val = nil
	} else {
		val = v.Value[0].Eval()
	}
	consts.SYMBOL_TABLE[v.Ident.Key] = val
	return val
}

func (n *Var) CompileJs(b *strings.Builder) {
	if n.Ident == nil {
		return
	}
	// js does not want let for already declared variables
	if _, ok := consts.SYMBOL_TABLE[n.Ident.Key]; !ok {
		consts.SYMBOL_TABLE[n.Ident.Key] = true
		b.WriteString("let ")
	}
	n.Ident.CompileJs(b)
	if len(n.Value) > 1 {
		b.WriteString("=")
		b.WriteRune('[')
		for i, c := range n.Value {
			if c == nil {
				continue
			}
			c.CompileJs(b)
			if i+1 < len(n.Value) {
				b.WriteRune(',')
			}
		}
		b.WriteRune(']')
	} else if len(n.Value) == 1 {
		v := n.Value[0]
		if v == nil {
			return
		}
		b.WriteString("=")
		v.CompileJs(b)
	}
}
