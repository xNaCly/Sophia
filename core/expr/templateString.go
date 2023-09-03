package expr

import (
	"fmt"
	"sophia/core/token"
	"strings"
)

type TemplateString struct {
	Token    token.Token
	Children []Node
}

func (s *TemplateString) GetToken() token.Token {
	return s.Token
}

func (s *TemplateString) Eval() any {
	b := strings.Builder{}
	for _, c := range s.Children {
		b.WriteString(fmt.Sprint(c.Eval()))
	}
	return b.String()
}
func (n *TemplateString) CompileJs(b *strings.Builder) {
	b.WriteRune('`')
	for _, c := range n.Children {
		t := c.GetToken()
		if t.Type == token.STRING {
			b.WriteString(t.Raw)
		} else if t.Type == token.IDENT {
			b.WriteString("${")
			b.WriteString(t.Raw)
			b.WriteRune('}')
		}
	}
	b.WriteRune('`')
}
