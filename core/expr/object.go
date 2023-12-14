package expr

import (
	"sophia/core/serror"
	"sophia/core/token"
	"strings"
)

type ObjectPair struct {
	Key   Node
	Value Node
}

type Object struct {
	Token    *token.Token
	Children []ObjectPair
}

func (o *Object) GetChildren() []Node {
	return nil
}

func (n *Object) SetChildren(c []Node) {}

func (o *Object) GetToken() *token.Token {
	return o.Token
}

func (o *Object) Eval() any {
	m := make(map[string]any, len(o.Children))
	for _, c := range o.Children {
		ident, ok := c.Key.(*Ident)
		if !ok {
			t := c.Key.GetToken()
			// TODO: support floats as object keys? idk should i?
			serror.Add(t, "Illegal object key", "Can not use %q as object key, use any identifier", token.TOKEN_NAME_MAP[t.Type])
			serror.Panic()
		}
		m[ident.Name] = c.Value.Eval()
	}
	return m
}

func (o *Object) CompileJs(b *strings.Builder) {
	b.WriteRune('{')
	for i, c := range o.Children {
		c.Key.CompileJs(b)
		b.WriteRune(':')
		c.Value.CompileJs(b)
		if i+1 < len(o.Children) {
			b.WriteRune(',')
		}
	}
	b.WriteRune('}')
}
