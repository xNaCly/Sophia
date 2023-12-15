package expr

import (
	"sophia/core/consts"
	"sophia/core/token"
	"sophia/core/types"
)

// defining a variable
type Var struct {
	Token *token.Token
	Ident *Ident
	Value []types.Node
}

func (v *Var) GetChildren() []types.Node {
	return v.Value
}

func (n *Var) SetChildren(c []types.Node) {
	n.Value = c
}

func (v *Var) GetToken() *token.Token {
	return v.Token
}

func (v *Var) Eval() any {
	var val any
	if len(v.Value) > 1 {
		val = make([]any, len(v.Value))
		for i, c := range v.Value {
			val.([]any)[i] = c.Eval()
		}
	} else if len(v.Value) == 0 {
		val = nil
	} else {
		val = v.Value[0].Eval()
	}
	consts.SYMBOL_TABLE[v.Ident.Key] = val
	return val
}
