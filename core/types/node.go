package types

import (
	"github.com/xnacly/sophia/core/token"
)

type Node interface {
	GetToken() *token.Token
	GetChildren() []Node
	SetChildren(c []Node)
	Eval() any
}
