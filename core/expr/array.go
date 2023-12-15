package expr

import (
	"sophia/core/token"
	"sophia/core/types"
)

type Array struct {
	Token    *token.Token
	Children []types.Node
}

func (a *Array) GetChildren() []types.Node {
	return a.Children
}

func (a *Array) SetChildren(c []types.Node) {
	a.Children = c
}

func (a *Array) GetToken() *token.Token {
	return a.Token
}

func (a *Array) Eval() any {
	if len(a.Children) == 0 {
		return []any{}
	}

	m := make([]any, 0, len(a.Children))
	for i := 0; i < len(a.Children); i++ {
		m = append(m, a.Children[i].Eval())
	}

	return m
}
