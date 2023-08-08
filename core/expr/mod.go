package expr

import "sophia/core/token"

type Mod struct {
	Token    token.Token
	Children []Node
}

func (m *Mod) GetToken() token.Token {
	return m.Token
}

func (m *Mod) Eval() any {
	if len(m.Children) == 0 {
		return 0.0
	}

	res := extractChild(m.Children[0], token.MOD)
	for _, c := range m.Children[1:] {
		res = float64(int(res) % int(extractChild(c, token.MOD)))
	}
	return res
}
