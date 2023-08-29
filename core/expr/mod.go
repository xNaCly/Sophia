package expr

import (
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
func (n *Mod) CompileJs(b *strings.Builder) {}
