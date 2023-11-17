package expr

import (
	"bytes"
	"sophia/core/token"
	"strings"
)

var templateBuffer = &bytes.Buffer{}

type TemplateString struct {
	Token    *token.Token
	Children []Node
}

func (s *TemplateString) GetChildren() []Node {
	return s.Children
}

func (n *TemplateString) SetChildren(c []Node) {
	n.Children = c
}

func (s *TemplateString) GetToken() *token.Token {
	return s.Token
}

func (s *TemplateString) Eval() any {
	if len(s.Children) == 0 {
		return ""
	}

	templateBuffer.Reset()
	formatHelper(buffer, s.Children)
	return templateBuffer.String()
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
