package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type Or struct {
	Token    token.Token
	Children []Node
}

func (o *Or) GetToken() token.Token {
	return o.Token
}

func (o *Or) Eval() any {
	for _, c := range o.Children {
		if castPanicIfNotType[bool](c.Eval(), token.OR) {
			return true
		}
	}
	return false
}
func (n *Or) CompileJs(b *strings.Builder) {
	cLen := len(n.Children)
	if cLen == 1 || cLen == 0 {
		debug.Log("opt: replaced illogical 'or' expression containing one or less children with true at line", n.Token.Line)
		b.WriteString("true")
		return
	}
	for i, c := range n.Children {
		c.CompileJs(b)
		if i+1 < cLen {
			b.WriteString("||")
		}
	}
}
