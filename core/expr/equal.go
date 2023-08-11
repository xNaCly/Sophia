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
	list := make([]any, len(e.Children))
	for i, c := range e.Children {
		list[i] = c.Eval()
		if i >= 1 && list[i-1] != list[i] {
			return false
		}
	}
	return true
}
