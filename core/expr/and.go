package expr

import (
	"sophia/core/token"
	"strings"
)

type And struct {
	Token    *token.Token
	Children []Node
}

func (a *And) GetChildren() []Node {
	return a.Children
}

func (n *And) SetChildren(c []Node) {
	n.Children = c
}

func (a *And) GetToken() *token.Token {
	return a.Token
}

func (a *And) Eval() any {
	// fastpaths
	if len(a.Children) == 2 {
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
	for i, c := range n.Children {
		c.CompileJs(b)
		if i+1 < cLen {
			b.WriteString("&&")
		}
	}
}
