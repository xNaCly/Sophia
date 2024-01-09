package expr

import (
	"github.com/xnacly/sophia/core/consts"
	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
)

// defining a variable
type Var struct {
	Token       *token.Token
	IndexAssign bool
	Ident       *Ident
	Value       []types.Node
}

func (v *Var) GetChildren() []types.Node {
	return v.Value
}

func (n *Var) SetChildren(c []types.Node) {
	n.Value = c
}

func (v *Var) GetToken() *token.Token {
	return v.Token
}

func (v *Var) Eval() any {
	var val any
	if len(v.Value) > 1 {
		val = make([]any, len(v.Value))
		for i, c := range v.Value {
			val.([]any)[i] = c.Eval()
		}
	} else if len(v.Value) == 0 {
		val = nil
	} else {
		// TODO: implement assignment to object and array

		// example:
		// ;; vim: syntax=lisp
		// (let tracker {})
		// (let names "anon" "anon1" "anon" "anon")
		// (for (_ name) names
		//     (let tracker[name] 1)
		//     (println name))
		//
		// (println tracker)

		if v.IndexAssign {
			serror.Add(v.Ident.Token, "Not implemented", "Assignment to array or object is currently not implemented - sorry :(")
			serror.Panic()
		}
		val = v.Value[0].Eval()
	}

	consts.SYMBOL_TABLE[v.Ident.Key] = val
	return val
}
