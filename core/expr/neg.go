package expr

import (
	"sophia/core/token"
	"strings"
)

type Neg struct {
	Token    token.Token
	Children Node
}

func (n *Neg) GetToken() token.Token {
	return n.Token
}

func (n *Neg) Eval() any {
	return !castPanicIfNotType[bool](n.Children.Eval(), n.Children.GetToken())
}
func (n *Neg) CompileJs(b *strings.Builder) {
	b.WriteRune('!')
	n.Children.CompileJs(b)
}
