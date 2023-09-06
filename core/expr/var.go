package expr

import (
	"sophia/core/consts"
	"sophia/core/token"
	"strings"
)

// defining a variable
type Var struct {
	Token token.Token
	Name  string
	Value []Node
}

func (v *Var) GetToken() token.Token {
	return v.Token
}

func (v *Var) Eval() any {
	var val any
	if len(v.Value) > 1 {
		val = make([]any, len(v.Value))
		for i, c := range v.Value {
			val.([]any)[i] = c.Eval()
		}
	} else if len(v.Value) == 0 {
		val = nil
	} else {
		val = v.Value[0].Eval()
	}

	consts.SYMBOL_TABLE[v.Name] = val
	return val
}
func (n *Var) CompileJs(b *strings.Builder) {
	b.WriteString("let ")
	b.WriteString(n.Name)
	if len(n.Value) > 1 {
		b.WriteString("=")
		b.WriteRune('[')
		for i, c := range n.Value {
			c.CompileJs(b)
			if i+1 < len(n.Value) {
				b.WriteRune(',')
			}
		}
		b.WriteRune(']')
	} else if len(n.Value) == 1 {
		b.WriteString("=")
		n.Value[0].CompileJs(b)
	}
}
