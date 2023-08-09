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
		ev := c.Eval()
		if val, ok := isType[[]interface{}](ev); ok {
			for _, s := range val {
				b.WriteString(castPanicIfNotType[string](s, token.CONCAT))
			}
		} else {
			b.WriteString(castPanicIfNotType[string](ev, token.CONCAT))
		}
	}
	return b.String()
}
