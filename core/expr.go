package core

type Node interface {
	GetToken() Token
}

type Statement struct {
	Token    Token
	Children []Node
}

func (p *Statement) GetToken() Token {
	return p.Token
}

type Float struct {
	Token Token
}

func (f *Float) GetToken() Token {
	return f.Token
}

type String struct {
	Token Token
}

func (s *String) GetToken() Token {
	return s.Token
}

type Putv struct {
	Token    Token
	Children []Node
}
