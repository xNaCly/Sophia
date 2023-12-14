package expr

import (
	"sophia/core/consts"
	"sophia/core/token"
"sophia/core/types"
	"strings"
)

// function definition
type Func struct {
	Token  *token.Token
	Name   types.Node
	Params types.Node
	Body   []types.Node
}

func (f *Func) GetChildren() []types.Node {
	return f.Body
}

func (n *Func) SetChildren(c []types.Node) {
	n.Body = c
}

func (f *Func) GetToken() *token.Token {
	return f.Token
}

func (f *Func) Eval() any {
	ident := f.Name.(*Ident)
	consts.FUNC_TABLE[ident.Key] = f
	return nil
}

func (n *Func) CompileJs(b *strings.Builder) {
	cLen := len(n.Body)
	b.WriteString("function ")
	b.WriteString(n.Name.GetToken().Raw)
	b.WriteRune('(')
	n.Params.CompileJs(b)
	b.WriteString("){")
	for i, c := range n.Body {
		if i+1 == cLen {
			b.WriteString("return ")
		}
		c.CompileJs(b)
		b.WriteRune(';')
	}
	b.WriteRune('}')
}
