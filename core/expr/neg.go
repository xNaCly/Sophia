package expr

import (
	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
)

type Neg struct {
	Token    *token.Token
	Children types.Node
}

func (n *Neg) GetChildren() []types.Node {
	return []types.Node{n.Children}
}

func (n *Neg) SetChildren(c []types.Node) {
	if len(c) == 0 {
		return
	}
	n.Children = c[0]
}

func (n *Neg) GetToken() *token.Token {
	return n.Token
}

func (n *Neg) Eval() any {
	child := n.Children.Eval()
	switch v := child.(type) {
	case nil:
		return false
	case float64:
		return v * -1
	case bool:
		return !v
	default:
		t := n.Children.GetToken()
		serror.Add(t, "Type Error", "Expected float64, bool or nil, got %T", child)
		serror.Panic()
	}
	return nil
}
