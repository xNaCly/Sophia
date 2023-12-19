package expr

import (
	"sophia/core/token"
	"sophia/core/types"
)

type Any struct {
	Value any
}

func (a *Any) GetChildren() []types.Node {
	return nil
}

func (a *Any) SetChildren(c []types.Node) {
}

func (a *Any) GetToken() *token.Token {
	return nil
}

func (a *Any) Eval() any {
	return a.Value
}
