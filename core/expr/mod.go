package expr

import (
	"math"
	"sophia/core/token"
"sophia/core/types"
	"strings"
)

type Mod struct {
	Token    *token.Token
	Children []types.Node
}

func (m *Mod) GetChildren() []types.Node {
	return m.Children
}

func (n *Mod) SetChildren(c []types.Node) {
	n.Children = c
}

func (m *Mod) GetToken() *token.Token {
	return m.Token
}

func (m *Mod) Eval() any {
	if len(m.Children) == 2 {
		// fastpath for two children
		f := m.Children[0]
		s := m.Children[1]
		return math.Mod(castFloatPanic(f.Eval(), f.GetToken()), castFloatPanic(s.Eval(), s.GetToken()))
	}

	res := 0.0
	for i, c := range m.Children {
		if i == 0 {
			res = castFloatPanic(c.Eval(), c.GetToken())
		} else {
			res = math.Mod(res, castFloatPanic(c.Eval(), c.GetToken()))
		}
	}
	return float64(res)
}
func (n *Mod) CompileJs(b *strings.Builder) {
	for i, c := range n.Children {
		c.CompileJs(b)
		if i+1 < len(n.Children) {
			b.WriteRune('%')
		}
	}
}
