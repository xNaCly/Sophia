package expr

import (
	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
)

type Module struct {
	Token    *token.Token
	Children []types.Node
}

func (m *Module) GetChildren() []types.Node {
	return m.Children
}

func (m *Module) SetChildren(c []types.Node) {
	m.Children = c
}

func (m *Module) GetToken() *token.Token {
	return m.Token
}

func (m *Module) Eval() any {
	serror.Add(m.Token, "Not implemented", "Modules are a work in progress.")
	serror.Panic()
	return nil
}
