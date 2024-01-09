package eval

import (
	"fmt"
	"github.com/xnacly/sophia/core/types"
)

func Eval(t string, ast []types.Node) []string {
	if t == "repl" {
		r := make([]string, len(ast))
		for i, c := range ast {
			r[i] = fmt.Sprint(c.Eval())
		}
		return r
	}
	for _, c := range ast {
		c.Eval()
	}
	return []string{}
}
