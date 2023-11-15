package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type And struct {
	Token    *token.Token
	Children []Node
}

func (a *And) GetToken() *token.Token {
	return a.Token
}

func (a *And) Eval() any {
	// fastpaths
	if len(a.Children) == 0 {
		return false
	} else if len(a.Children) == 1 {
		c := a.Children[0]
		return castBoolPanic(c.Eval(), c.GetToken())
	} else if len(a.Children) == 2 {
		f := a.Children[0]
		s := a.Children[1]
		return castBoolPanic(f.Eval(), f.GetToken()) && castBoolPanic(s.Eval(), s.GetToken())
	}

	for _, c := range a.Children {
		v := castBoolPanic(c.Eval(), a.Token)
		if !v {
			return false
		}
	}
	return true
}

func (n *And) CompileJs(b *strings.Builder) {
	cLen := len(n.Children)
	if cLen == 1 || cLen == 0 {
		debug.Log("opt: replaced illogical 'and' expression containing one or less children with true at line", n.Token.Line)
		b.WriteString("true")
		return
	}
	for i, c := range n.Children {
		c.CompileJs(b)
		if i+1 < cLen {
			b.WriteString("&&")
		}
	}
}
