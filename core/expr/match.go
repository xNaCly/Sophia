package expr

import (
	"sophia/core/token"
	"strings"
)

type Match struct {
	Token    *token.Token
	Branches []Node
}

func (m *Match) GetToken() *token.Token {
	return m.Token
}

func (m *Match) Eval() any {
	// fastpath: skip loop and lookup
	if len(m.Branches) == 0 {
		return nil
	}
	for _, c := range m.Branches {
		if c.GetToken().Type == token.IF {
			o := c.Eval()
			if o.(bool) {
				return nil
			}
		} else {
			return c.Eval()
		}
	}
	return nil
}

func (n *Match) CompileJs(b *strings.Builder) {
	if len(n.Branches) == 1 {
		n.Branches[0].CompileJs(b)
	}
	for i, c := range n.Branches {
		isIf := c.GetToken().Type != token.IF
		if i != 0 {
			b.WriteString("else")
		}
		if isIf {
			b.WriteRune('{')
		} else {
			b.WriteRune(' ')
		}
		c.CompileJs(b)
		if c.GetToken().Type != token.IF {
			b.WriteRune('}')
		}
	}
}
