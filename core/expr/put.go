package expr

import (
	"bytes"
	"os"
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

var buffer = &bytes.Buffer{}

type Put struct {
	Token    *token.Token
	Children []Node
}

func (p *Put) GetChildren() []Node {
	return p.Children
}

func (n *Put) SetChildren(c []Node) {
	n.Children = c
}

func (p *Put) GetToken() *token.Token {
	return p.Token
}

func (p *Put) Eval() any {
	buffer.Reset()
	formatHelper(buffer, p.Children, ' ')
	buffer.WriteRune('\n')
	buffer.WriteTo(os.Stdout)
	return nil
}

func (n *Put) CompileJs(b *strings.Builder) {
	cLen := len(n.Children)
	if cLen == 0 {
		debug.Log("opt: removed empty print")
		return
	}
	b.WriteString("console.log(")
	for i, c := range n.Children {
		c.CompileJs(b)
		if i+1 < cLen {
			b.WriteRune(',')
		}
	}
	b.WriteString(")")
}
