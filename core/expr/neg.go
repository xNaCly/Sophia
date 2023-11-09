package expr

import (
	"sophia/core/serror"
	"sophia/core/token"
	"strings"
)

type Neg struct {
	Token    token.Token
	Children Node
}

func (n *Neg) GetToken() token.Token {
	return n.Token
}

func (n *Neg) Eval() any {
	child := n.Children.Eval()
	var r any
	switch child.(type) {
	case float64:
		r = child.(float64) * -1
	case bool:
		r = !child.(bool)
	default:
		t := n.Children.GetToken()
		serror.Add(&t, "Type Error", "Expected float64 or bool, got %T", child)
		serror.Panic()
	}
	return r
}

func (n *Neg) CompileJs(b *strings.Builder) {
	b.WriteRune('!')
	n.Children.CompileJs(b)
}
