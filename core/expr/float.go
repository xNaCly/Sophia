package expr

import (
	"fmt"
	"sophia/core/token"
	"strconv"
)

type Float struct {
	Token token.Token
}

func (f *Float) GetToken() token.Token {
	return f.Token
}

func (f *Float) Eval() any {
	float, err := strconv.ParseFloat(f.Token.Raw, 64)
	if err != nil {
		panic(fmt.Sprint("failed to parse float: ", err))
	}
	return float
}
