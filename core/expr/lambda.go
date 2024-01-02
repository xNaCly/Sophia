package expr

import (
	"sophia/core/serror"
	"sophia/core/token"
	"sophia/core/types"
)

type Lambda struct {
	Token  *token.Token
	Body   []types.Node
	Params *Array
	Args   []types.Node
}

func (l *Lambda) GetChildren() []types.Node {
	return l.Args
}

func (l *Lambda) SetChildren(c []types.Node) {
	l.Args = c
}

func (l *Lambda) GetToken() *token.Token {
	return l.Token
}

func (l *Lambda) Eval() any {
	if len(l.Args) == 0 {
		serror.Add(l.Token, "Illogical lambda", "Lambda got no argument, consider using it with the map or filter built-ins")
		serror.Panic()
	}
	return callFunction(l.Token, l.Body, l.Params, l.Args)
}
