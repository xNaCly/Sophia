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
	res := 0.0
	for i, c := range m.Children {
		if i == 0 {
			res = castPanicIfNotType[float64](c.Eval(), token.MUL)
		} else {
			res *= castPanicIfNotType[float64](c.Eval(), token.MUL)
		}
	}
	return res
}
