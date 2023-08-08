package expr

import (
	"sophia/core/consts"
	"sophia/core/token"
)

// defining a variable
type Var struct {
	Token token.Token
	Name  string
	Value []Node
}

func (v *Var) GetToken() token.Token {
	return v.Token
}

func (v *Var) Eval() any {
	val := make([]any, len(v.Value))
	for i, c := range v.Value {
		val[i] = c.Eval()
	}

	if _, ok := consts.SYMBOL_TABLE[v.Name]; ok {
		consts.SYMBOL_TABLE[v.Name] = val
	} else {
		consts.SYMBOL_TABLE[v.Name] = val
	}
	return val
}
