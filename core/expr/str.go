package expr

import (
	"sophia/core/token"
	"strings"
)

type String struct {
	Token *token.Token
}

func (s *String) GetToken() *token.Token {
	return s.Token
}

func (s *String) Eval() any {
	return s.Token.Raw
}
func (n *String) CompileJs(b *strings.Builder) {
	b.WriteRune('"')
	b.WriteString(n.Token.Raw)
	b.WriteRune('"')
}
