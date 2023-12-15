package expr

import (
	"sophia/core/token"
	"sophia/core/types"
)

type Equal struct {
	Token    *token.Token
	Children []types.Node
}

func (e *Equal) GetChildren() []types.Node {
	return e.Children
}

func (n *Equal) SetChildren(c []types.Node) {
	n.Children = c
}

func (e *Equal) GetToken() *token.Token {
	return e.Token
}

func (e *Equal) Eval() any {
	if len(e.Children) == 2 {
		// skipping list creating for multiple equal children
		return e.Children[0].Eval() == e.Children[1].Eval()
	}
	list := make([]any, len(e.Children))
	for i, c := range e.Children {
		list[i] = c.Eval()
		if i >= 1 && list[i-1] != list[i] {
			return false
		}
	}
	return true
}
