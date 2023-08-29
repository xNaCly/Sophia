package expr

import (
	"fmt"
	"sophia/core/token"
	"strings"
)

type Put struct {
	Token    token.Token
	Children []Node
}

func (p *Put) GetToken() token.Token {
	return p.Token
}

func (p *Put) Eval() any {
	b := strings.Builder{}
	for i, c := range p.Children {
		if i != 0 {
			b.WriteRune(' ')
		}
		b.WriteString(fmt.Sprint(c.Eval()))
	}
	fmt.Printf("%s\n", b.String())
	return nil
}
func (n *Put) CompileJs(b *strings.Builder) {
	b.WriteString("console.log(")
	for i, c := range n.Children {
		c.CompileJs(b)
		if i+1 < len(n.Children) {
			b.WriteRune(',')
		}
	}
	b.WriteString(")")
}
