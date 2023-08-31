package eval

import (
	"fmt"
	"sophia/core/expr"
	"strings"
)

func Eval(ast []expr.Node) []string {
	out := make([]string, len(ast))
	for i, c := range ast {
		out[i] = fmt.Sprint(c.Eval())
	}
	return out
}

func CompileJs(ast []expr.Node) string {
	b := strings.Builder{}
	for _, c := range ast {
		c.CompileJs(&b)
		b.WriteRune(';')
	}
	return b.String()
}
