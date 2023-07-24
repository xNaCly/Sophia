package core

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func checkIfType[T any](in any, op int) (T, bool) {
	val, ok := in.(T)
	if !ok {
		var e T
		log.Printf("can not use variable of type %T in current operation (%s), expected %T for value %q", in, TOKEN_NAME_MAP[op], e, in)
		return e, false
	}
	return val, true
}

type Node interface {
	GetToken() Token
	Eval() any
}

type Statement struct {
	Token    Token
	Children []Node
}

func (s *Statement) GetToken() Token {
	return s.Token
}

func (s *Statement) Eval() any {
	for _, c := range s.Children {
		c.Eval()
	}
	return 0.0
}

type Float struct {
	Token Token
}

func (f *Float) GetToken() Token {
	return f.Token
}

func (f *Float) Eval() any {
	return f.Token.Float
}

type String struct {
	Token Token
}

func (s *String) GetToken() Token {
	return s.Token
}

func (s *String) Eval() any {
	return s.Token.Raw
}

type Put struct {
	Token    Token
	Children []Node
}

func (p *Put) GetToken() Token {
	return p.Token
}

func (p *Put) Eval() any {
	b := strings.Builder{}
	for i, c := range p.Children {
		if i != 0 {
			b.WriteRune(' ')
		}
		switch c.(type) {
		case *String:
			b.WriteString(c.GetToken().Raw)
			continue
		}
		b.WriteString(fmt.Sprint(c.Eval()))
	}
	fmt.Printf("%s\n", b.String())
	return 0.0
}

type Add struct {
	Token    Token
	Children []Node
}

func (a *Add) GetToken() Token {
	return a.Token
}

func (a *Add) Eval() any {
	if len(a.Children) == 0 {
		return 0.0
	}
	res, ok := checkIfType[float64](a.Children[0].Eval(), ADD)
	if !ok {
		return 0.0
	}
	for _, c := range a.Children[1:] {
		val, ok := checkIfType[float64](c.Eval(), ADD)
		if !ok {
			return 0.0
		}
		res += val
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

func (s *Sub) Eval() any {
	if len(s.Children) == 0 {
		return 0.0
	}
	res, ok := checkIfType[float64](s.Children[0].Eval(), SUB)
	if !ok {
		return 0.0
	}
	for _, c := range s.Children[1:] {
		val, ok := checkIfType[float64](c.Eval(), SUB)
		if !ok {
			return 0.0
		}
		res -= val
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

func (m *Mul) Eval() any {
	if len(m.Children) == 0 {
		return 0.0
	}
	res, ok := checkIfType[float64](m.Children[0].Eval(), MUL)
	if !ok {
		return 0.0
	}
	for _, c := range m.Children[1:] {
		val, ok := checkIfType[float64](c.Eval(), MUL)
		if !ok {
			return 0.0
		}
		res *= val
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

func (d *Div) Eval() any {
	if len(d.Children) == 0 {
		return 0.0
	}
	res, ok := checkIfType[float64](d.Children[0].Eval(), DIV)
	if !ok {
		return 0.0
	}
	for _, c := range d.Children[1:] {
		val, ok := checkIfType[float64](c.Eval(), DIV)
		if !ok {
			return 0.0
		}
		res /= val
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

func (p *Pwr) Eval() any {
	if len(p.Children) == 0 {
		return 0.0
	}
	res, ok := checkIfType[float64](p.Children[0].Eval(), PWR)
	if !ok {
		return 0.0
	}
	for _, c := range p.Children[1:] {
		val, ok := checkIfType[float64](c.Eval(), PWR)
		if !ok {
			return 0.0
		}
		res += math.Pow(res, val)
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

func (m *Mod) Eval() any {
	if len(m.Children) == 0 {
		return 0.0
	}
	r, ok := checkIfType[float64](m.Children[0].Eval(), MOD)
	if !ok {
		return 0.0
	}
	res := int(r)
	for _, c := range m.Children[1:] {
		val, ok := checkIfType[float64](c.Eval(), MOD)
		if !ok {
			return 0.0
		}
		res = res % int(val)
	}
	return float64(res)
}

type Ident struct {
	Token Token
	Name  string
}

func (i *Ident) GetToken() Token {
	return i.Token
}

func (i *Ident) Eval() any {
	val, ok := SYMBOL_TABLE[i.Name]
	if !ok {
		log.Printf("variable '%s' is not defined!", i.Name)
	}
	return val
}

type Var struct {
	Token Token
	Name  string
	Value Node
}

func (v *Var) GetToken() Token {
	return v.Token
}

func (v *Var) Eval() any {
	val := v.Value.Eval()
	if _, ok := SYMBOL_TABLE[v.Name]; ok {
		SYMBOL_TABLE[v.Name] = val
	} else {
		SYMBOL_TABLE[v.Name] = val
	}
	return val
}
