package expr

import (
	"sophia/core/token"
	"sophia/core/types"
)

type If struct {
	Token     *token.Token
	Condition types.Node
	Body      []types.Node
}

func (i *If) GetChildren() []types.Node {
	return i.Body
}

func (n *If) SetChildren(c []types.Node) {
	n.Body = c
}

func (i *If) GetToken() *token.Token {
	return i.Token
}

func (i *If) Eval() any {
	cond := castBoolPanic(i.Condition.Eval(), i.Condition.GetToken())
	if !cond {
		return false
	}
	for _, c := range i.Body {
		c.Eval()
	}
	return true
}
