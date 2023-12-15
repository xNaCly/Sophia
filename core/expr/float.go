package expr

import (
	"sophia/core/token"
	"sophia/core/types"
)

type Float struct {
	Token *token.Token
	Value float64
}

func (f *Float) GetChildren() []types.Node {
	return nil
}

func (n *Float) SetChildren(c []types.Node) {}

func (f *Float) GetToken() *token.Token {
	return f.Token
}

func (f *Float) Eval() any {
	return f.Value
}
