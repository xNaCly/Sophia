package expr

import (
	"sophia/core/token"
	"strings"
)

type Boolean struct {
	Token token.Token
}

func (b *Boolean) GetToken() token.Token {
	return b.Token
}

func (b *Boolean) Eval() any {
	return b.Token.Raw == "true"
}

func (n *Boolean) CompileJs(b *strings.Builder) {
	b.WriteString(n.Token.Raw)
}
