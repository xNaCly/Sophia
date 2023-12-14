package expr

import (
	"sophia/core/token"
"sophia/core/types"
	"strings"
)

type Gt struct {
	Token    *token.Token
	Children []types.Node
}

func (g *Gt) GetChildren() []types.Node {
	return g.Children
}

func (n *Gt) SetChildren(c []types.Node) {
	n.Children = c
}

func (g *Gt) GetToken() *token.Token {
	return g.Token
}

func (g *Gt) Eval() any {
	return castFloatPanic(g.Children[0].Eval(), g.Children[0].GetToken()) > castFloatPanic(g.Children[1].Eval(), g.Children[1].GetToken())
}

func (n *Gt) CompileJs(b *strings.Builder) {
	n.Children[0].CompileJs(b)
	b.WriteRune('<')
	n.Children[1].CompileJs(b)
}
