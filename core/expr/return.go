package expr

import (
	"sophia/core/consts"
	"sophia/core/token"
	"sophia/core/types"
)

type Return struct {
	Token *token.Token
	Child types.Node
}

func (r *Return) GetChildren() []types.Node {
	return []types.Node{r.Child}
}

func (r *Return) SetChildren(c []types.Node) {
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
