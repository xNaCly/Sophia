package expr

import (
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
func (n *Div) CompileJs(b *strings.Builder) {}
