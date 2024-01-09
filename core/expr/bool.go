package expr

import (
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
)

type Boolean struct {
	Token *token.Token
	Value bool
}

func (b *Boolean) GetChildren() []types.Node {
	return nil
}

func (n *Boolean) SetChildren(c []types.Node) {}

func (b *Boolean) GetToken() *token.Token {
	return b.Token
}

func (b *Boolean) Eval() any {
	return b.Value
}
