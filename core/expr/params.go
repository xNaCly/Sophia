package expr

import (
	"sophia/core/token"
	"strings"
)

type Params struct {
	Token    token.Token
	Children []Node
}

func (p *Params) GetToken() token.Token {
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
