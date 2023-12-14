package eval

import (
	"fmt"
	"sophia/core/types"
	"strings"
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

func CompileJs(ast []types.Node) string {
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
