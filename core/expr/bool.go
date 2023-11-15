package expr

import (
	"sophia/core/token"
	"strings"
)

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) GetToken() token.Token {
	return b.Token
}

func (b *Boolean) Eval() any {
	return b.Value
}

func (n *Boolean) CompileJs(b *strings.Builder) {
	b.WriteString(n.Token.Raw)
}
