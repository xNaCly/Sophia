package expr

import (
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
)

type String struct {
	Token *token.Token
}

func (s *String) GetChildren() []types.Node {
	return nil
}

func (n *String) SetChildren(c []types.Node) {}

func (s *String) GetToken() *token.Token {
	return s.Token
}

func (s *String) Eval() any {
	return s.Token.Raw
}
