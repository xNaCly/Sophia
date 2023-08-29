package expr

import (
	"sophia/core/token"
	"strings"
)

type If struct {
	Token     token.Token
	Condition Node
	Body      []Node
}

func (i *If) GetToken() token.Token {
	return i.Token
}

func (i *If) Eval() any {
	cond := castPanicIfNotType[bool](i.Condition.Eval(), token.IF)
	if cond {
		for _, c := range i.Body {
			c.Eval()
		}
	}
	return cond
}
func (n *If) CompileJs(b *strings.Builder) {}
