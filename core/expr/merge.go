package expr

import (
	"sophia/core/token"
	"strings"
)

type Merge struct {
	Token    token.Token
	Children []Node
}

func (m *Merge) GetToken() token.Token {
	return m.Token
}

func (m *Merge) Eval() any {
	if len(m.Children) == 0 {
		return nil
	}

	evaledChilds := make([]interface{}, len(m.Children))
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
func (n *Merge) CompileJs(b *strings.Builder) {}
