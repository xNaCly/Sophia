package expr

import (
	"fmt"
	"sophia/core/consts"
	"sophia/core/token"
)

// function definition
type For struct {
	Token    token.Token
	Params   Node
	LoopOver Node
	Body     []Node
}

func (f *For) GetToken() token.Token {
	return f.Token
}

func (f *For) Eval() any {
	params := f.Params.(*Params).Children
	if len(params) < 1 {
		panic(fmt.Sprintf("expected at least %d parameters in loop, got %d", 1, len(params)))
	}
	element := castPanicIfNotType[*Ident](params[0], token.FOR)
	oldValue, foundOldValue := consts.SYMBOL_TABLE[element.Name]

	loopOver := castPanicIfNotType[[]interface{}](f.LoopOver.Eval(), token.FOR)

	for _, el := range loopOver {
		consts.SYMBOL_TABLE[element.Name] = el
		for _, stmt := range f.Body {
			stmt.Eval()
		}
	}

	defer func() {
		if foundOldValue {
			consts.SYMBOL_TABLE[element.Name] = oldValue
		}
	}()
	return nil
}
