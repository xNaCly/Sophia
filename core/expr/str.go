package expr

import "sophia/core/token"

type String struct {
	Token token.Token
}

func (s *String) GetToken() token.Token {
	return s.Token
}

func (s *String) Eval() any {
	return s.Token.Raw
}
