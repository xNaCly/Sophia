package eval

import (
	"fmt"
	"sophia/core/expr"
)

func Eval(ast []expr.Node) []string {
	out := make([]string, len(ast))
	for i, c := range ast {
		out[i] = fmt.Sprint(c.Eval())
	}
	fmt.Println(out)
	return out
}
