package eval

import (
	"fmt"
	"sophia/core/expr"
	"strings"
)

func Eval(t string, ast []expr.Node) []string {
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

func CompileJs(ast []expr.Node) string {
	b := strings.Builder{}
	for _, c := range ast {
		l := b.Len()

		c.CompileJs(&b)

		// INFO: if compiling the last expression did yield a result, append a
		// semicolon, because we are at the top level expression
		if b.Len() != l {
			b.WriteRune(';')
		}
	}
	return b.String()
}
