package expr

import (
	"sophia/core/token"
	"strings"
)

type Gt struct {
	Token    token.Token
	Children []Node
}

func (g *Gt) GetToken() token.Token {
	return g.Token
}

func (g *Gt) Eval() any {
	return castFloatPanic(g.Children[0].Eval(), g.Children[0].GetToken()) > castFloatPanic(g.Children[1].Eval(), g.Children[1].GetToken())
}

func (n *Gt) CompileJs(b *strings.Builder) {
	b.WriteString(n.Children[0].GetToken().Raw)
	b.WriteRune('>')
	b.WriteString(n.Children[1].GetToken().Raw)
}
