package expr

import (
	"sophia/core/debug"
	"sophia/core/token"
	"strings"
)

type Add struct {
	Token    *token.Token
	Children []Node
}

func (a *Add) GetToken() *token.Token {
	return a.Token
}

func (a *Add) Eval() any {
	if len(a.Children) == 0 {
		return 0.0
	} else if len(a.Children) == 1 {
		// fastpath for skipping loop and casts
		return a.Children[0].Eval()
	} else if len(a.Children) == 2 {
		// fastpath for two children
		f := a.Children[0]
		s := a.Children[1]
		return castFloatPanic(f.Eval(), f.GetToken()) + castFloatPanic(s.Eval(), s.GetToken())
	}

	res := 0.0
	for i, c := range a.Children {
		if i == 0 {
			res = castFloatPanic(c.Eval(), a.Token)
		} else {
			res += castFloatPanic(c.Eval(), a.Token)
		}
	}
	return res
}

func (n *Add) CompileJs(b *strings.Builder) {
	cLen := len(n.Children)
	if cLen == 0 {
		debug.Log("opt: removed illogical '+' expression containing no children at line", n.Token.Line)
	} else {
		for i, c := range n.Children {
			c.CompileJs(b)
			if i+1 < cLen {
				b.WriteRune('+')
			}
		}
	}
}
