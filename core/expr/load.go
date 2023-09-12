package expr

import (
	"sophia/core/token"
	"strings"
)

type Load struct {
	Token   token.Token
	Imports []string
}

func (l *Load) GetToken() token.Token {
	return l.Token
}

func (l *Load) Eval() any {
	return nil
}
func (n *Load) CompileJs(b *strings.Builder) {}