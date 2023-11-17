package expr

import (
	"sophia/core/token"
	"strings"
)

type If struct {
	Token     *token.Token
	Condition Node
	Body      []Node
}

func (i *If) GetChildren() []Node {
	return i.Body
}

func (n *If) SetChildren(c []Node) {
	n.Body = c
}

func (i *If) GetToken() *token.Token {
	return i.Token
}

func (i *If) Eval() any {
	cond := castBoolPanic(i.Condition.Eval(), i.Condition.GetToken())
	if !cond {
		return false
	}
	for _, c := range i.Body {
		c.Eval()
	}
	return true
}
func (n *If) CompileJs(b *strings.Builder) {
	b.WriteString("if(")
	n.Condition.CompileJs(b)
	b.WriteRune(')')
	b.WriteRune('{')
	for _, c := range n.Body {
		c.CompileJs(b)
		b.WriteRune(';')
	}
	b.WriteRune('}')
}
