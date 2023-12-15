package expr

import (
	"sophia/core/consts"
	"sophia/core/serror"
	"sophia/core/token"
	"sophia/core/types"
)

// using a variable
type Ident struct {
	Token *token.Token
	Key   uint32
	Name  string
}

func (i *Ident) GetChildren() []types.Node {
	return []types.Node{}
}

func (n *Ident) SetChildren(c []types.Node) {}

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
