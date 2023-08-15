package expr

import "sophia/core/token"

type Gt struct {
	Token    token.Token
	Children []Node
}

func (g *Gt) GetToken() token.Token {
	return g.Token
}

func (g *Gt) Eval() any {
	return castPanicIfNotType[float64](g.Children[0].Eval(), token.GT) > castPanicIfNotType[float64](g.Children[1].Eval(), token.GT)
}
