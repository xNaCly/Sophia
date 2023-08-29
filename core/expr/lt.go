package expr

import (
	"sophia/core/token"
	"strings"
)

type Lt struct {
	Token    token.Token
	Children []Node
}

func (l *Lt) GetToken() token.Token {
	return l.Token
}

func (l *Lt) Eval() any {
	return castPanicIfNotType[float64](l.Children[0].Eval(), token.LT) < castPanicIfNotType[float64](l.Children[1].Eval(), token.LT)
}
func (n *Lt) CompileJs(b *strings.Builder) {}
