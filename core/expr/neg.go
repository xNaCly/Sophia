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
	ev := n.Children.Eval()
	return !castPanicIfNotType[bool](ev, token.NEG)
}
func (n *Neg) CompileJs(b *strings.Builder) {}
