package expr

import (
	"sophia/core/token"
	"strings"
)

// parser only structure
type Load struct {
	Token   *token.Token
	Imports []Node
}

func (l *Load) GetChildren() []Node {
	return nil
}

func (n *Load) SetChildren(c []Node) {}

func (l *Load) GetToken() *token.Token {
	return l.Token
}

func (l *Load) Eval() any {
	return nil
}

func (n *Load) CompileJs(b *strings.Builder) {}
