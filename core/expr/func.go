package expr

import (
	"sophia/core/consts"
	"sophia/core/token"
	"strings"
)

// function definition
type Func struct {
	Token  token.Token
	Name   Node
	Params Node
	Body   []Node
}

func (f *Func) GetToken() token.Token {
	return f.Token
}

func (f *Func) Eval() any {
	consts.FUNC_TABLE[f.Name.GetToken().Raw] = f
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
