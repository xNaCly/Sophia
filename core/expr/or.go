package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type Or struct {
	Token    *token.Token
	Children []Node
}

func (o *Or) GetChildren() []Node {
	return o.Children
}

func (n *Or) SetChildren(c []Node) {
	n.Children = c
}

func (o *Or) GetToken() *token.Token {
	return o.Token
}

func (o *Or) Eval() any {
	if len(o.Children) == 2 {
		f := o.Children[0]
		s := o.Children[1]
		return castBoolPanic(f.Eval(), f.GetToken()) || castBoolPanic(s.Eval(), s.GetToken())
	}
	for _, c := range o.Children {
		if castBoolPanic(c.Eval(), c.GetToken()) {
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
