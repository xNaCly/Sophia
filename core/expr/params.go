package expr

import "sophia/core/token"

type Params struct {
	Token    token.Token
	Children []Node
}

func (p *Params) GetToken() token.Token {
	return p.Token
}

func (p *Params) Eval() any {
	// TODO: add params to symbol table, somehow :(
	return nil
}
