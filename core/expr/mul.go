package expr

import (
	"sophia/core/token"
	"strings"
)

type Mul struct {
	Token    *token.Token
	Children []Node
}

func (m *Mul) GetChildren() []Node {
	return m.Children
}

func (n *Mul) SetChildren(c []Node) {
	n.Children = c
}

func (m *Mul) GetToken() *token.Token {
	return m.Token
}

func (m *Mul) Eval() any {
	if len(m.Children) == 2 {
		// fastpath for two children
		f := m.Children[0]
		s := m.Children[1]
		return castFloatPanic(f.Eval(), f.GetToken()) * castFloatPanic(s.Eval(), s.GetToken())
	}

	res := 0.0
	for i, c := range m.Children {
		if i == 0 {
			res = castFloatPanic(c.Eval(), c.GetToken())
		} else {
			res *= castFloatPanic(c.Eval(), c.GetToken())
		}
	}
	return res
}
func (n *Mul) CompileJs(b *strings.Builder) {
	for i, c := range n.Children {
		c.CompileJs(b)
		if i+1 < len(n.Children) {
			b.WriteRune('*')
		}
	}
}
