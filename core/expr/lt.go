package expr

import (
	"sophia/core/token"
	"strings"
)

type Lt struct {
	Token    *token.Token
	Children []Node
}

func (l *Lt) GetToken() *token.Token {
	return l.Token
}

func (l *Lt) Eval() any {
	return castFloatPanic(l.Children[0].Eval(), l.Children[0].GetToken()) < castFloatPanic(l.Children[1].Eval(), l.Children[1].GetToken())
}
func (n *Lt) CompileJs(b *strings.Builder) {
	b.WriteString(n.Children[0].GetToken().Raw)
	b.WriteRune('<')
	b.WriteString(n.Children[1].GetToken().Raw)
}
