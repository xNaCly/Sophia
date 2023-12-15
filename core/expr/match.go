package expr

import (
	"sophia/core/token"
	"sophia/core/types"
)

type Match struct {
	Token    *token.Token
	Branches []types.Node
}

func (m *Match) GetChildren() []types.Node {
	return m.Branches
}

func (n *Match) SetChildren(c []types.Node) {
	n.Branches = c
}

func (m *Match) GetToken() *token.Token {
	return m.Token
}

func (m *Match) Eval() any {
	// fastpath: skip loop and lookup
	if len(m.Branches) == 0 {
		return nil
	}
	for _, c := range m.Branches {
		if c.GetToken().Type == token.IF {
			o := c.Eval()
			if o.(bool) {
				return nil
			}
		} else {
			return c.Eval()
		}
	}
	return nil
}
