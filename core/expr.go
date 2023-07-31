package core

import (
	"fmt"
	"strings"
)

// TODO: error display

// attempts to cast `in` to `T`, returns `in` cast to `T` if successful. If
// cast fails, panics.
func castPanicIfNotType[T any](in any, op int) T {
	val, ok := isType[T](in)
	if !ok {
		panic(fmt.Sprintf("can not use variable of type %T in current operation (%s), expected %T for value %+v", in, TOKEN_NAME_MAP[op], val, in))
	}
	return val
}

func isType[T any](in any) (T, bool) {
	val, ok := in.(T)
	if !ok {
		var e T
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
	// TODO: assign the first value of the children to this
	res := 0.0
	for _, c := range a.Children[1:] {
		if idt, ok := isType[*Ident](c); ok {
			arr := castPanicIfNotType[[]interface{}](idt.Eval(), ADD)
			for _, i := range arr {
				res += castPanicIfNotType[float64](i, ADD)
			}
		} else {
			res += castPanicIfNotType[float64](c.Eval(), ADD)
		}
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
	// TODO: assign the first value of the children to this
	res := 0.0
	for _, c := range s.Children {
		if idt, ok := isType[*Ident](c); ok {
			arr := castPanicIfNotType[[]interface{}](idt.Eval(), SUB)
			for _, i := range arr {
				res -= castPanicIfNotType[float64](i, SUB)
			}
		} else {
			res -= castPanicIfNotType[float64](c.Eval(), SUB)
		}
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
	// TODO: assign the first value of the children to this
	res := 1.0
	for _, c := range m.Children {
		if idt, ok := isType[*Ident](c); ok {
			arr := castPanicIfNotType[[]interface{}](idt.Eval(), MUL)
			for _, i := range arr {
				res *= castPanicIfNotType[float64](i, MUL)
			}
		} else {
			res *= castPanicIfNotType[float64](c.Eval(), MUL)
		}
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
	// TODO: assign the first value of the children to this
	res := 0.0
	for _, c := range d.Children {
		if idt, ok := isType[*Ident](c); ok {
			arr := castPanicIfNotType[[]interface{}](idt.Eval(), DIV)
			for _, i := range arr {
				res /= castPanicIfNotType[float64](i, DIV)
			}
		} else {
			res /= castPanicIfNotType[float64](c.Eval(), DIV)
		}
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
	// TODO:
	res := 0
	for _, c := range m.Children {
		if idt, ok := isType[*Ident](c); ok {
			arr := castPanicIfNotType[[]interface{}](idt.Eval(), MOD)
			for _, i := range arr {
				res = res % int(castPanicIfNotType[float64](i, MOD))
			}
		} else {
			res = res % int(castPanicIfNotType[float64](c.Eval(), MOD))
		}
	}
	return res
}

// using a variable
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
		panic(fmt.Sprintf("variable '%s' is not defined!", i.Name))
	}
	return val
}

// defining a variable
type Var struct {
	Token Token
	Name  string
	Value []Node
}

func (v *Var) GetToken() Token {
	return v.Token
}

func (v *Var) Eval() any {
	val := make([]any, len(v.Value))
	for i, c := range v.Value {
		val[i] = c.Eval()
	}

	if _, ok := SYMBOL_TABLE[v.Name]; ok {
		SYMBOL_TABLE[v.Name] = val
	} else {
		SYMBOL_TABLE[v.Name] = val
	}
	return val
}
