package expr

import (
	"sophia/core/token"
	"strings"
)

type Array struct {
	Token    *token.Token
	Children []Node
}

func (a *Array) GetChildren() []Node {
	return a.Children
}

func (a *Array) SetChildren(c []Node) {
	a.Children = c
}

func (a *Array) GetToken() *token.Token {
	return a.Token
}

func (a *Array) Eval() any {
	if len(a.Children) == 0 {
		return []any{}
	}

	m := make([]any, 0, len(a.Children))
	for i := 0; i < len(a.Children); i++ {
		m = append(m, a.Children[i].Eval())
	}

	return m
}

func (a *Array) CompileJs(b *strings.Builder) {
	b.WriteRune('[')
	for i, c := range a.Children {
		c.CompileJs(b)
		if i+1 != len(a.Children) {
			b.WriteRune(',')
		}
	}
	b.WriteRune(']')
}
