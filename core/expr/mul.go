package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type Mul struct {
	Token    *token.Token
	Children []Node
}

func (m *Mul) GetToken() *token.Token {
	return m.Token
}

func (m *Mul) Eval() any {
	if len(m.Children) == 0 {
		return 0.0
	} else if len(m.Children) == 1 {
		// fastpath for skipping loop and casts
		return m.Children[0].Eval()
	} else if len(m.Children) == 2 {
		// fastpath for two children
		f := m.Children[0]
		s := m.Children[1]
		return castFloatPanic(f.Eval(), f.GetToken()) * castFloatPanic(s.Eval(), s.GetToken())
	}

	res := 0.0
	for i, c := range m.Children {
		if i == 0 {
			res = castFloatPanic(c.Eval(), m.Token)
		} else {
			res *= castFloatPanic(c.Eval(), m.Token)
		}
	}
	return res
}
func (n *Mul) CompileJs(b *strings.Builder) {
	cLen := len(n.Children)
	if cLen == 0 {
		debug.Log("opt: removed illogical '*' expression containing zero children at line", n.Token.Line)
		return
	} else {
		for i, c := range n.Children {
			c.CompileJs(b)
			if i+1 < cLen {
				b.WriteRune('*')
			}
		}
	}
}
