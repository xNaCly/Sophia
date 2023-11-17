package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type Div struct {
	Token    *token.Token
	Children []Node
}

func (d *Div) GetToken() *token.Token {
	return d.Token
}

func (d *Div) Eval() any {
	if len(d.Children) == 0 {
		return 0.0
	} else if len(d.Children) == 1 {
		// fastpath for skipping loop and casts
		return d.Children[0].Eval()
	} else if len(d.Children) == 2 {
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
	cLen := len(n.Children)
	if cLen == 0 {
		debug.Log("opt: removed illogical '/' expression containing no children at line", n.Token.Line)
		return
	} else {
		for i, c := range n.Children {
			c.CompileJs(b)
			if i+1 < cLen {
				b.WriteRune('/')
			}
		}
	}
}
