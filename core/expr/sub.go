package expr

import "sophia/core/token"

type Sub struct {
	Token    token.Token
	Children []Node
}

func (s *Sub) GetToken() token.Token {
	return s.Token
}

func (s *Sub) Eval() any {
	if len(s.Children) == 0 {
		return 0.0
	}
	res := extractChild(s.Children[0], token.SUB)
	for _, c := range s.Children[1:] {
		res -= extractChild(c, token.SUB)
	}
	return res
}
