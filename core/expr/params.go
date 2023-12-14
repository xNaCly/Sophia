package expr

import (
	"sophia/core/token"
"sophia/core/types"
	"strings"
)

type Params struct {
	Token    *token.Token
	Children []types.Node
}

func (p *Params) GetChildren() []types.Node {
	return p.Children
}

func (n *Params) SetChildren(c []types.Node) {
	n.Children = c
}

func (p *Params) GetToken() *token.Token {
	return p.Token
}

func (p *Params) Eval() any {
	return nil
}

func (n *Params) CompileJs(b *strings.Builder) {
	for i, c := range n.Children {
		c.CompileJs(b)
		if i+1 < len(n.Children) {
			b.WriteRune(',')
		}
	}
}
