package expr

import "sophia/core/token"

type Boolean struct {
	Token token.Token
}

func (b *Boolean) GetToken() token.Token {
	return b.Token
}

func (b *Boolean) Eval() any {
	return b.Token.Raw == "true"
}
