package expr

import (
	"sophia/core/consts"
	"sophia/core/token"
	"strings"
)

type Return struct {
	Token *token.Token
	Child Node
}

func (r *Return) GetChildren() []Node {
	return []Node{r.Child}
}

func (r *Return) SetChildren(c []Node) {
	r.Child = c[0]
}

func (r *Return) GetToken() *token.Token {
	return r.Token
}

func (r *Return) Eval() any {
	if r.Child == nil {
		return nil
	}
	e := r.Child.Eval()
	consts.RETURN.HasValue = true
	consts.RETURN.Value = e
	return e
}

func (r *Return) CompileJs(b *strings.Builder) {
	b.WriteString("return ")
	r.Child.CompileJs(b)
}
