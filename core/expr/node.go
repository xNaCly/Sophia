package expr

import (
	"sophia/core/token"
	"strings"
)

type Node interface {
	GetToken() *token.Token
	GetChildren() []Node
	SetChildren(c []Node)
	Eval() any
	CompileJs(b *strings.Builder)
}
