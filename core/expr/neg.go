package expr

import (
	"sophia/core/serror"
	"sophia/core/token"
	"strings"
)

type Neg struct {
	Token    *token.Token
	Children Node
}

func (n *Neg) GetToken() *token.Token {
	return n.Token
}

func (n *Neg) Eval() any {
	child := n.Children.Eval()
	var r any
	switch v := child.(type) {
	case float64:
		r = v * -1
	case bool:
		r = !v
	default:
		t := n.Children.GetToken()
		serror.Add(t, "Type Error", "Expected float64 or bool, got %T", child)
		serror.Panic()
	}
	return r
}

func (n *Neg) CompileJs(b *strings.Builder) {
	switch n.Children.(type) {
	case *Float:
		b.WriteRune('-')
	case *Boolean:
		b.WriteRune('!')
	}
	n.Children.CompileJs(b)
}
