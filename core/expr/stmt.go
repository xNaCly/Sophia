package expr

import "sophia/core/token"

type Statement struct {
	Token    token.Token
	Children []Node
}

func (s *Statement) GetToken() token.Token {
	return s.Token
}

func (s *Statement) Eval() any {
	for _, c := range s.Children {
		c.Eval()
	}
	return 0.0
}
