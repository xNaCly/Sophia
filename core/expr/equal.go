package expr

import "sophia/core/token"

type Equal struct {
	Token    token.Token
	Children []Node
}

func (e *Equal) GetToken() token.Token {
	return e.Token
}

func (e *Equal) Eval() any {
	list := make([]any, 0)
	for _, c := range e.Children {
		ev := c.Eval()
		if val, ok := isType[[]interface{}](ev); ok {
			list = append(list, val...)
		} else {
			list = append(list, ev)
		}
	}
	for i := range list {
		if i > 0 {
			if list[i-1] != list[i] {
				return false
			}
		}
	}
	return true
}
