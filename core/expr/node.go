package expr

import (
	"sophia/core/token"
	"strings"
)

type Node interface {
	GetToken() *token.Token
	Eval() any
	CompileJs(b *strings.Builder)
}
