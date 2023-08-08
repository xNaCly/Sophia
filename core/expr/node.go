package expr

import "sophia/core/token"

type Node interface {
	GetToken() token.Token
	Eval() any
}
