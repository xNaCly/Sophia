package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
"sophia/core/types"
	"strings"
)

type Merge struct {
	Token    *token.Token
	Children []types.Node
}

func (m *Merge) GetChildren() []types.Node {
	return m.Children
}

func (n *Merge) SetChildren(c []types.Node) {
	n.Children = c
}

func (m *Merge) GetToken() *token.Token {
	return m.Token
}

func (m *Merge) Eval() any {
	if len(m.Children) == 1 {
		return []any{m.Children[0].Eval()}
	}

	evaledChilds := make([]any, len(m.Children))
	tryString := true
	for i, c := range m.Children {
		evaledChilds[i] = c.Eval()
		if _, ok := evaledChilds[i].(string); !ok {
			tryString = false
		}
	}

	if tryString {
		if val, ok := evaledChilds[0].(string); ok {
			b := strings.Builder{}
			b.WriteString(val)
			for _, c := range evaledChilds[1:] {
				if out, ok := c.(string); ok {
					b.WriteString(out)
				}
			}
			return b.String()
		}
	}

	merged := make([]interface{}, 0)
	for _, el := range evaledChilds {
		if val, ok := el.([]interface{}); ok {
			merged = append(merged, val...)
		} else {
			merged = append(merged, el)
		}
	}
	return merged
}

func (n *Merge) CompileJs(b *strings.Builder) {
	cLen := len(n.Children)
	if cLen == 0 {
		debug.Log("opt: removed illogical '++' expression containing no children on line ", n.Token.Line)
		return
	}
	b.WriteString("[].concat(")
	for i, c := range n.Children {
		c.CompileJs(b)
		if i+1 < cLen {
			b.WriteRune(',')
		}
	}
	b.WriteRune(')')
}
