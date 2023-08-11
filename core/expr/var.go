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
	var val any
	if len(v.Value) > 1 {
		val = make([]any, len(v.Value))
		for i, c := range v.Value {
			val.([]any)[i] = c.Eval()
		}
	} else {
		val = v.Value[0].Eval()
	}

	consts.SYMBOL_TABLE[v.Name] = val
	return val
}
