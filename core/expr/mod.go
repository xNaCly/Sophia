package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type Mod struct {
	Token    token.Token
	Children []Node
}

func (m *Mod) GetToken() token.Token {
	return m.Token
}

func (m *Mod) Eval() any {
	if len(m.Children) == 0 {
		return 0.0
	}

	res := 0
	for i, c := range m.Children {
		if i == 0 {
			res = int(castPanicIfNotType[float64](c.Eval(), token.MOD))
		} else {
			res = res % int(castPanicIfNotType[float64](c.Eval(), token.MOD))
		}
	}
	return float64(res)
}
func (n *Mod) CompileJs(b *strings.Builder) {
	cLen := len(n.Children)
	if cLen == 0 {
		debug.Log("opt: removed illogical '%' expression containing no children at line", n.Token.Line)
		return
	} else {
		for i, c := range n.Children {
			c.CompileJs(b)
			if i+1 < cLen {
				b.WriteRune('%')
			}
		}
	}
}
