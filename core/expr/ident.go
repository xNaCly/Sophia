package expr

import (
	"fmt"
	"sophia/core/consts"
	"sophia/core/token"
)

// using a variable
type Ident struct {
	Token token.Token
	Name  string
}

func (i *Ident) GetToken() token.Token {
	return i.Token
}

func (i *Ident) Eval() any {
	val, ok := consts.SYMBOL_TABLE[i.Name]
	if !ok {
		panic(fmt.Sprintf("variable '%s' is not defined!", i.Name))
	}
	return val
}
