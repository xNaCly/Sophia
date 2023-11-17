package expr

import (
	"sophia/core/token"
	"strings"
)

type Add struct {
	Token    *token.Token
	Children []Node
}

func (a *Add) GetChildren() []Node {
	return a.Children
}

func (n *Add) SetChildren(c []Node) {
	n.Children = c
}

func (a *Add) GetToken() *token.Token {
	return a.Token
}

func (a *Add) Eval() any {
	if len(a.Children) == 2 {
		// fastpath for two children
		f := a.Children[0]
		s := a.Children[1]
		return castFloatPanic(f.Eval(), f.GetToken()) + castFloatPanic(s.Eval(), s.GetToken())
	}

	res := 0.0
	for i, c := range a.Children {
		if i == 0 {
			res = castFloatPanic(c.Eval(), c.GetToken())
		} else {
			res += castFloatPanic(c.Eval(), c.GetToken())
		}
	}
	return res
}

func (n *Add) CompileJs(b *strings.Builder) {
	for i, c := range n.Children {
		c.CompileJs(b)
		if i+1 < len(n.Children) {
			b.WriteRune('+')
		}
	}
}
