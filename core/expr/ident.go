package expr

import (
	"sophia/core/consts"
	"sophia/core/serror"
	"sophia/core/token"
	"strings"
)

// using a variable
type Ident struct {
	Token *token.Token
	Name  string
}

func (i *Ident) GetChildren() []Node {
	return nil
}

func (n *Ident) SetChildren(c []Node) {}

func (i *Ident) GetToken() *token.Token {
	return i.Token
}

func (i *Ident) Eval() any {
	val, ok := consts.SYMBOL_TABLE[i.Name]
	if !ok {
		serror.Add(i.Token, "Undefined variable", "Variable %q is not defined.", i.Name)
		serror.Panic()
	}
	return val
}

func (n *Ident) CompileJs(b *strings.Builder) {
	b.WriteString(n.Name)
}
