package core

import (
	"fmt"
	"math"
)

type Node interface {
	GetToken() Token
	Eval() float64
}

type Statement struct {
	Token    Token
	Children []Node
}

func (s *Statement) GetToken() Token {
	return s.Token
}

func (s *Statement) Eval() float64 {
	for _, c := range s.Children {
		c.Eval()
	}
	return 0
}

type Float struct {
	Token Token
}

func (f *Float) GetToken() Token {
	return f.Token
}

func (f *Float) Eval() float64 {
	return f.Token.Float
}

type String struct {
	Token Token
}

func (s *String) GetToken() Token {
	return s.Token
}

func (s *String) Eval() float64 {
	return 0
}

type Put struct {
	Token    Token
	Children []Node
}

func (p *Put) GetToken() Token {
	return p.Token
}

func (p *Put) Eval() float64 {
	res := make([]any, 0)
	for _, c := range p.Children {
		switch c.(type) {
		case *String:
			res = append(res, c.GetToken().Raw)
			continue
		}
		res = append(res, c.Eval())
	}
	fmt.Printf("~ %v\n", res)
	return 0
}

type Add struct {
	Token    Token
	Children []Node
}

func (a *Add) GetToken() Token {
	return a.Token
}

func (a *Add) Eval() float64 {
	if len(a.Children) == 0 {
		return 0
	}
	res := a.Children[0].Eval()
	for _, c := range a.Children[1:] {
		res += c.Eval()
	}
	return res
}

type Sub struct {
	Token    Token
	Children []Node
}

func (s *Sub) GetToken() Token {
	return s.Token
}

func (s *Sub) Eval() float64 {
	if len(s.Children) == 0 {
		return 0
	}
	res := s.Children[0].Eval()
	for _, c := range s.Children[1:] {
		res -= c.Eval()
	}
	return res
}

type Mul struct {
	Token    Token
	Children []Node
}

func (m *Mul) GetToken() Token {
	return m.Token
}

func (m *Mul) Eval() float64 {
	if len(m.Children) == 0 {
		return 0
	}
	res := m.Children[0].Eval()
	for _, c := range m.Children[1:] {
		res *= c.Eval()
	}
	return res
}

type Div struct {
	Token    Token
	Children []Node
}

func (d *Div) GetToken() Token {
	return d.Token
}

func (d *Div) Eval() float64 {
	if len(d.Children) == 0 {
		return 0
	}
	res := d.Children[0].Eval()
	for _, c := range d.Children[1:] {
		res /= c.Eval()
	}
	return res
}

type Pwr struct {
	Token    Token
	Children []Node
}

func (p *Pwr) GetToken() Token {
	return p.Token
}

func (p *Pwr) Eval() float64 {
	if len(p.Children) == 0 {
		return 0
	}
	res := p.Children[0].Eval()
	for _, c := range p.Children[1:] {
		res = math.Pow(res, c.Eval())
	}
	return res
}

type Mod struct {
	Token    Token
	Children []Node
}

func (m *Mod) GetToken() Token {
	return m.Token
}

func (m *Mod) Eval() float64 {
	if len(m.Children) == 0 {
		return 0
	}
	res := int(m.Children[0].Eval())
	for _, c := range m.Children[1:] {
		res = res % int(c.Eval())
	}
	return float64(res)
}
