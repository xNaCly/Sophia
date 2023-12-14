package expr

import (
	"sophia/core/token"
"sophia/core/types"
	"strings"
)

// parser only structure
type Load struct {
	Token   *token.Token
	Imports []types.Node
}

func (l *Load) GetChildren() []types.Node {
	return nil
}

func (n *Load) SetChildren(c []types.Node) {}

func (l *Load) GetToken() *token.Token {
	return l.Token
}

func (l *Load) Eval() any {
	return nil
}

func (n *Load) CompileJs(b *strings.Builder) {}
