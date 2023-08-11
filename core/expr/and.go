package expr

import "sophia/core/token"

type And struct {
	Token    token.Token
	Children []Node
}

func (a *And) GetToken() token.Token {
	return a.Token
}

func (a *And) Eval() any {
	for _, c := range a.Children {
		v := castPanicIfNotType[bool](c.Eval(), token.AND)
		if !v {
			return false
		}
	}
	return true
}
