package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type Div struct {
	Token    token.Token
	Children []Node
}

func (d *Div) GetToken() token.Token {
	return d.Token
}

func (d *Div) Eval() any {
	if len(d.Children) == 0 {
		return 0.0
	}
	res := 0.0
	for i, c := range d.Children {
		if i == 0 {
			res = castPanicIfNotType[float64](c.Eval(), token.DIV)
		} else {
			res /= castPanicIfNotType[float64](c.Eval(), token.DIV)
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
