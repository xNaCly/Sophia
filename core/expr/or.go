package expr

import "sophia/core/token"

type Or struct {
	Token    token.Token
	Children []Node
}

func (o *Or) GetToken() token.Token {
	return o.Token
}

func (o *Or) Eval() any {
	list := make([]bool, 0)
	for _, c := range o.Children {
		ev := c.Eval()
		if val, ok := isType[[]interface{}](ev); ok {
			for _, v := range val {
				list = append(list, castPanicIfNotType[bool](v, token.OR))
			}
		} else {
			list = append(list, castPanicIfNotType[bool](ev, token.OR))
		}
	}
	for _, v := range list {
		if v {
			return true
		}
	}
	return false
}
