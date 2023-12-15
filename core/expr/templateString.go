package expr

import (
	"sophia/core/shared"
	"sophia/core/token"
	"sophia/core/types"
	"strings"
)

var buffer = &strings.Builder{}

type TemplateString struct {
	Token    *token.Token
	Children []types.Node
}

func (s *TemplateString) GetChildren() []types.Node {
	return s.Children
}

func (n *TemplateString) SetChildren(c []types.Node) {
	n.Children = c
}

func (s *TemplateString) GetToken() *token.Token {
	return s.Token
}

func (s *TemplateString) Eval() any {
	if len(s.Children) == 0 {
		return ""
	}

	buffer.Reset()
	shared.FormatHelper(buffer, s.Children, 0)
	return buffer.String()
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
