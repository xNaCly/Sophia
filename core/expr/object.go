package expr

import (
	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
)

type ObjectPair struct {
	Key   types.Node
	Value types.Node
}

type Object struct {
	Token    *token.Token
	Children []ObjectPair
}

func (o *Object) GetChildren() []types.Node {
	return nil
}

func (n *Object) SetChildren(c []types.Node) {}

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
