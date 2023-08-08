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
	list := make([]bool, 0)
	for _, c := range a.Children {
		ev := c.Eval()
		if val, ok := isType[[]interface{}](ev); ok {
			for _, v := range val {
				list = append(list, castPanicIfNotType[bool](v, token.AND))
			}
		} else {
			list = append(list, castPanicIfNotType[bool](ev, token.AND))
		}
	}
	for _, v := range list {
		if !v {
			return false
		}
	}
	return true
}
