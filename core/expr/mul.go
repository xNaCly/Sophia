package expr

import "sophia/core/token"

type Mul struct {
	Token    token.Token
	Children []Node
}

func (m *Mul) GetToken() token.Token {
	return m.Token
}

func (m *Mul) Eval() any {
	if len(m.Children) == 0 {
		return 0.0
	}
	res := extractChild(m.Children[0], token.MUL)
	for _, c := range m.Children[1:] {
		res *= extractChild(c, token.MUL)
	}
	return res
}
