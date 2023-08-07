package core

import (
	"fmt"
	"strconv"
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

// checks if `in` is castable to `T`, returns casted value and true if
// castable, zero value of `T` and false if not
func isType[T any](in any) (T, bool) {
	val, ok := in.(T)
	if !ok {
		var e T
		return e, false
	}
	return val, true
}

func extractChild(n Node, op int) float64 {
	var val float64
	if idt, ok := isType[*Ident](n); ok {
		arr := castPanicIfNotType[[]interface{}](idt.Eval(), op)
		for i, item := range arr {
			t := castPanicIfNotType[float64](item, op)
			if i == 0 {
				val = t
				continue
			}
			switch op {
			case ADD:
				val += t
			case SUB:
				val -= t
			case DIV:
				val /= t
			case MUL:
				val *= t
			case MOD:
				val = float64(int(val) % int(t))
			}
		}
	} else {
		val = castPanicIfNotType[float64](n.Eval(), op)
	}
	return val
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
	float, err := strconv.ParseFloat(f.Token.Raw, 64)
	if err != nil {
		panic(fmt.Sprint("failed to parse float: ", err))
	}
	return float
}

type Boolean struct {
	Token Token
}

func (b *Boolean) GetToken() Token {
	return b.Token
}

func (b *Boolean) Eval() any {
	return b.Token.Raw == "true"
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
	return nil
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
	res := extractChild(a.Children[0], ADD)
	for _, c := range a.Children[1:] {
		res += extractChild(c, ADD)
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
	res := extractChild(s.Children[0], SUB)
	for _, c := range s.Children[1:] {
		res -= extractChild(c, SUB)
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
	res := extractChild(m.Children[0], MUL)
	for _, c := range m.Children[1:] {
		res *= extractChild(c, MUL)
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
	res := extractChild(d.Children[0], DIV)
	for _, c := range d.Children[1:] {
		res /= extractChild(c, DIV)
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

	res := extractChild(m.Children[0], MOD)
	for _, c := range m.Children[1:] {
		res = float64(int(res) % int(extractChild(c, MOD)))
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

type If struct {
	Token     Token
	Condition Node
	Body      []Node
}

func (i *If) GetToken() Token {
	return i.Token
}

func (i *If) Eval() any {
	if castPanicIfNotType[bool](i.Condition.Eval(), IF) {
		for _, c := range i.Body {
			c.Eval()
		}
	}
	return nil
}

type Equal struct {
	Token    Token
	Children []Node
}

func (e *Equal) GetToken() Token {
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

type Or struct {
	Token    Token
	Children []Node
}

func (o *Or) GetToken() Token {
	return o.Token
}

func (o *Or) Eval() any {
	list := make([]bool, 0)
	for _, c := range o.Children {
		ev := c.Eval()
		if val, ok := isType[[]interface{}](ev); ok {
			for _, v := range val {
				list = append(list, castPanicIfNotType[bool](v, OR))
			}
		} else {
			list = append(list, castPanicIfNotType[bool](ev, OR))
		}
	}
	for _, v := range list {
		if v {
			return true
		}
	}
	return false
}

type And struct {
	Token    Token
	Children []Node
}

func (a *And) GetToken() Token {
	return a.Token
}

func (a *And) Eval() any {
	list := make([]bool, 0)
	for _, c := range a.Children {
		ev := c.Eval()
		if val, ok := isType[[]interface{}](ev); ok {
			for _, v := range val {
				list = append(list, castPanicIfNotType[bool](v, AND))
			}
		} else {
			list = append(list, castPanicIfNotType[bool](ev, AND))
		}
	}
	for _, v := range list {
		if !v {
			return false
		}
	}
	return true
}

type Neg struct {
	Token    Token
	Children Node
}

func (n *Neg) GetToken() Token {
	return n.Token
}

func (n *Neg) Eval() any {
	ev := n.Children.Eval()
	return !castPanicIfNotType[bool](ev, NEG)
}
