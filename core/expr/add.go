package expr

import (
	"sophia/core/debug"
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

func (n *Add) CompileJs(b *strings.Builder) {
	cLen := len(n.Children)
	if cLen == 0 {
		debug.Log("opt: removed illogical '+' expression containing no children at line", n.Token.Line)
	} else {
		for i, c := range n.Children {
			c.CompileJs(b)
			if i+1 < cLen {
				b.WriteRune('+')
			}
		}
	}
}
