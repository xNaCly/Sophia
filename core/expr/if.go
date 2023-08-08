package expr

import "sophia/core/token"

type If struct {
	Token     token.Token
	Condition Node
	Body      []Node
}

func (i *If) GetToken() token.Token {
	return i.Token
}

func (i *If) Eval() any {
	if castPanicIfNotType[bool](i.Condition.Eval(), token.IF) {
		for _, c := range i.Body {
			c.Eval()
		}
	}
	return nil
}
