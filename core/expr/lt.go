package expr

import (
	"sophia/core/token"
"sophia/core/types"
	"strings"
)

type Lt struct {
	Token    *token.Token
	Children []types.Node
}

func (l *Lt) GetChildren() []types.Node {
	return l.Children
}

func (n *Lt) SetChildren(c []types.Node) {
	n.Children = c
}

func (l *Lt) GetToken() *token.Token {
	return l.Token
}

func (l *Lt) Eval() any {
	return castFloatPanic(l.Children[0].Eval(), l.Children[0].GetToken()) < castFloatPanic(l.Children[1].Eval(), l.Children[1].GetToken())
}

func (n *Lt) CompileJs(b *strings.Builder) {
	n.Children[0].CompileJs(b)
	b.WriteRune('<')
	n.Children[1].CompileJs(b)
}
