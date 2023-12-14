package expr

import (
	"sophia/core/token"
"sophia/core/types"
	"strings"
)

type Div struct {
	Token    *token.Token
	Children []types.Node
}

func (d *Div) GetChildren() []types.Node {
	return d.Children
}

func (n *Div) SetChildren(c []types.Node) {
	n.Children = c
}

func (d *Div) GetToken() *token.Token {
	return d.Token
}

func (d *Div) Eval() any {
	if len(d.Children) == 2 {
		// fastpath for two children
		f := d.Children[0]
		s := d.Children[1]
		return castFloatPanic(f.Eval(), f.GetToken()) / castFloatPanic(s.Eval(), s.GetToken())
	}
	res := 0.0
	for i, c := range d.Children {
		if i == 0 {
			res = castFloatPanic(c.Eval(), c.GetToken())
		} else {
			res /= castFloatPanic(c.Eval(), c.GetToken())
		}
	}
	return res
}
func (n *Div) CompileJs(b *strings.Builder) {
	for i, c := range n.Children {
		c.CompileJs(b)
		if i+1 < len(n.Children) {
			b.WriteRune('/')
		}
	}
}
