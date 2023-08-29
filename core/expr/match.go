package expr

import (
	"sophia/core/token"
	"strings"
)

type Match struct {
	Token    token.Token
	Branches []Node
}

func (m *Match) GetToken() token.Token {
	return m.Token
}

func (m *Match) Eval() any {
	for _, c := range m.Branches {
		if c.GetToken().Type == token.IF {
			if c.Eval().(bool) {
				break
			}
		} else {
			c.Eval()
		}
	}
	return nil
}
func (n *Match) CompileJs(b *strings.Builder) {}
