package expr

import (
	"fmt"
	"sophia/core/consts"
	"sophia/core/token"
)

type Call struct {
	Token  token.Token
	Params []Node
}

func (c *Call) GetToken() token.Token {
	return c.Token
}

func (c *Call) Eval() any {
	storedFunc, ok := consts.FUNC_TABLE[c.Token.Raw]
	if !ok {
		panic(fmt.Sprintf("function %q not defined", c.Token.Raw))
	}

	def := castPanicIfNotType[*Func](storedFunc, token.IDENT)
	defParams := castPanicIfNotType[*Params](def.Params, token.IDENT)

	if len(defParams.Children) != len(c.Params) {
		paramLen := len(defParams.Children)
		argLen := len(c.Params)
		if paramLen < argLen {
			panic(fmt.Sprintf("too many arguments for %q, wanted %d, got %d", c.Token.Raw, len(defParams.Children), len(c.Params)))
		} else if paramLen > argLen {
			panic(fmt.Sprintf("not enough arguments for %q, wanted %d, got %d", c.Token.Raw, len(defParams.Children), len(c.Params)))
		}
	}

	// TODO: this is not smart, this overrides variables in the global symbol
	// table, maybe prefix with hashes or the function name?
	for i, arg := range c.Params {
		consts.SYMBOL_TABLE[defParams.Children[i].GetToken().Raw] = arg.Eval()
	}

	// going out of scope, therefore we remove the parameter from the symbol table
	defer func() {
		for _, param := range defParams.Children {
			delete(consts.SYMBOL_TABLE, param.GetToken().Raw)
		}
	}()

	for i, stmt := range def.Body {
		if i+1 == len(def.Body) {
			return stmt.Eval()
		}
		stmt.Eval()
	}

	return nil
}
