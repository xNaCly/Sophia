package expr

import "sophia/core/token"

type Add struct {
	Token    token.Token
	Children []Node
}

func (a *Add) GetToken() token.Token {
	return a.Token
}

func (a *Add) Eval() any {
	if len(a.Children) == 0 {
		return 0.0
	}
	res := extractChild(a.Children[0], token.ADD)
	for _, c := range a.Children[1:] {
		res += extractChild(c, token.ADD)
	}
	return res
}
