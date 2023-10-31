package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type Sub struct {
	Token    token.Token
	Children []Node
}

func (s *Sub) GetToken() token.Token {
	return s.Token
}

func (s *Sub) Eval() any {
	if len(s.Children) == 0 {
		return 0.0
	}
	res := 0.0
	for i, c := range s.Children {
		if i == 0 {
			res = castPanicIfNotType[float64](c.Eval(), c.GetToken())
		} else {
			res -= castPanicIfNotType[float64](c.Eval(), c.GetToken())
		}
	}
	return res
}
func (n *Sub) CompileJs(b *strings.Builder) {
	cLen := len(n.Children)
	if cLen == 0 {
		debug.Log("opt: removed illogical '-' expression containing zero children at line", n.Token.Line)
	} else if cLen == 1 {
		b.WriteRune('-')
		n.Children[0].CompileJs(b)
	} else {
		for i, c := range n.Children {
			c.CompileJs(b)
			if i+1 < cLen {
				b.WriteRune('-')
			}
		}
	}
}
