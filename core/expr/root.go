package expr

import (
	"sophia/core/token"
	"strings"
)

type Root struct {
	Children []Node
}

func (r *Root) GetChildren() []Node {
	return r.Children
}

func (r *Root) SetChildren(c []Node) {
	r.Children = c
}

func (r *Root) GetToken() *token.Token {
	return nil
}

func (r *Root) Eval() any {
	return nil
}

func (r *Root) CompileJs(b *strings.Builder) {}
