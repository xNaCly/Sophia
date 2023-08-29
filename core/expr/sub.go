package expr

import (
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
			res = castPanicIfNotType[float64](c.Eval(), token.SUB)
		} else {
			res -= castPanicIfNotType[float64](c.Eval(), token.SUB)
		}
	}
	return res
}
func (n *Sub) CompileJs(b *strings.Builder) {}
