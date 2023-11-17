package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type If struct {
	Token     *token.Token
	Condition Node
	Body      []Node
}

func (i *If) GetToken() *token.Token {
	return i.Token
}

func (i *If) Eval() any {
	if len(i.Body) == 0 {
		debug.Log("opt: removed 'for loop' with no body at line", i.Token.Line)
		return false
	}
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
	if len(n.Body) == 0 {
		debug.Log("opt: removed 'if' with no body at line", n.Token.Line)
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
