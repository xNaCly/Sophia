package expr

import (
	"sophia/core/token"
	"strings"
)

type Add struct {
	Token    token.Token
	Children []Node
}

func (a *Add) GetToken() token.Token {
	return a.Token
}

func (a *Add) Eval() any {
	if len(a.Children) == 0 {
		return 0.0
	}
	res := 0.0
	for i, c := range a.Children {
		if i == 0 {
			res = castPanicIfNotType[float64](c.Eval(), token.ADD)
		} else {
			res += castPanicIfNotType[float64](c.Eval(), token.ADD)
		}
	}
	return res
}

func (n *Add) CompileJs(b *strings.Builder) {}
