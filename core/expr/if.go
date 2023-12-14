package expr

import (
	"sophia/core/token"
"sophia/core/types"
	"strings"
)

type If struct {
	Token     *token.Token
	Condition types.Node
	Body      []types.Node
}

func (i *If) GetChildren() []types.Node {
	return i.Body
}

func (n *If) SetChildren(c []types.Node) {
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
