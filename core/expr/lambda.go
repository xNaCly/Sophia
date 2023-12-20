package expr

// TODO: think about parsing, where should the arguments go if not used in map or filter, allow outside of filter or map?
// TODO: planned syntax: (lambda [params] (expression)*), useful for iterating arrays via map or filter:
// (let arr [1 2 3 4 5 6])
// (let arr
//      (filter
//          (lambda [n] (= (% n 2) 0)) ;; returns copy of arr containing all elements matching the lambda expression
//      )
// )

import (
	"sophia/core/token"
	"sophia/core/types"
)

type Lambda struct {
	Token     *token.Token
	Body      []types.Node
	Params    *Array
	Arguments []types.Node
}

func (l *Lambda) GetChildren() []types.Node {
	return l.Body
}

func (l *Lambda) SetChildren(c []types.Node) {
	l.Body = c
}

func (l *Lambda) GetToken() *token.Token {
	return l.Token
}

func (l *Lambda) Eval() any {
	return callFunction(l.Token, l.Body, l.Params, l.Arguments)
}
