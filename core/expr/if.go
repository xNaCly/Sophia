package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type If struct {
	Token     token.Token
	Condition Node
	Body      []Node
}

func (i *If) GetToken() token.Token {
	return i.Token
}

func (i *If) Eval() any {
	cond := castBoolPanic(i.Condition.Eval(), i.Condition.GetToken())
	if cond {
		for _, c := range i.Body {
			c.Eval()
		}
	}
	return cond
}
func (n *If) CompileJs(b *strings.Builder) {
	if len(n.Body) == 0 {
		debug.Log("opt: removed empty if at line", n.Token.Line)
		return
	}
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
