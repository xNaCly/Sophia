package expr

import "sophia/core/token"

type Div struct {
	Token    token.Token
	Children []Node
}

func (d *Div) GetToken() token.Token {
	return d.Token
}

func (d *Div) Eval() any {
	if len(d.Children) == 0 {
		return 0.0
	}
	res := extractChild(d.Children[0], token.DIV)
	for _, c := range d.Children[1:] {
		res /= extractChild(c, token.DIV)
	}
	return res
}
