package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type Mul struct {
	Token    token.Token
	Children []Node
}

func (m *Mul) GetToken() token.Token {
	return m.Token
}

func (m *Mul) Eval() any {
	if len(m.Children) == 0 {
		return 0.0
	}
	res := 0.0
	for i, c := range m.Children {
		if i == 0 {
			res = castPanicIfNotType[float64](c.Eval(), token.MUL)
		} else {
			res *= castPanicIfNotType[float64](c.Eval(), token.MUL)
		}
	}
	return res
}
func (n *Mul) CompileJs(b *strings.Builder) {
	cLen := len(n.Children)
	if cLen == 0 || cLen == 1 {
		debug.Log("opt: removed illogical '*' expression containing one or less children at line", n.Token.Line)
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
