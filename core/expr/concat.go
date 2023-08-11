package expr

import (
	"sophia/core/token"
	"strings"
)

type Concat struct {
	Token    token.Token
	Children []Node
}

func (c *Concat) GetToken() token.Token {
	return c.Token
}

func (c *Concat) Eval() any {
	b := strings.Builder{}
	for _, c := range c.Children {
		b.WriteString(castPanicIfNotType[string](c.Eval(), token.CONCAT))
	}
	return b.String()
}
