package expr

import (
	"sophia/core/token"
	"strings"
)

type Equal struct {
	Token    *token.Token
	Children []Node
}

func (e *Equal) GetChildren() []Node {
	return e.Children
}

func (n *Equal) SetChildren(c []Node) {
	n.Children = c
}

func (e *Equal) GetToken() *token.Token {
	return e.Token
}

func (e *Equal) Eval() any {
	if len(e.Children) == 2 {
		// skipping list creating for multiple equal children
		return e.Children[0].Eval() == e.Children[1].Eval()
	}
	list := make([]any, len(e.Children))
	for i, c := range e.Children {
		list[i] = c.Eval()
		if i >= 1 && list[i-1] != list[i] {
			return false
		}
	}
	return true
}

func (n *Equal) CompileJs(b *strings.Builder) {
	cLen := len(n.Children)
	if cLen == 2 {
		n.Children[0].CompileJs(b)
		b.WriteString("===")
		n.Children[1].CompileJs(b)
	} else {
		b.WriteRune('[')
		for i := 0; i < cLen; i++ {
			n.Children[i].CompileJs(b)
			if i+1 < cLen {
				b.WriteRune(',')
			}
		}
		b.WriteString("].every((e, i, arr) => e === arr[0])")
	}
}
