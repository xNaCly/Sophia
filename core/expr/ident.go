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
	Key   uint32
	Name  string
}

func (i *Ident) GetChildren() []Node {
	return []Node{}
}

func (n *Ident) SetChildren(c []Node) {}

func (i *Ident) GetToken() *token.Token {
	return i.Token
}

func (i *Ident) Eval() any {
	val, ok := consts.SYMBOL_TABLE[i.Key]
	if !ok {
		serror.Add(i.Token, "Undefined variable", "Variable %q is not defined.", i.Name)
		serror.Panic()
	}
	return val
}

func (n *Ident) CompileJs(b *strings.Builder) {
	b.WriteString(n.Name)
}
