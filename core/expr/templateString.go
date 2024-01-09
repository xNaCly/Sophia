package expr

import (
	"github.com/xnacly/sophia/core/shared"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
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
