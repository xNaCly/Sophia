package expr

import (
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
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
