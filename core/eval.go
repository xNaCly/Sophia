package core

import "fmt"

var SYMBOL_TABLE = map[string]any{}

func Eval(ast []Node) []string {
	out := make([]string, len(ast))
	for i, c := range ast {
		out[i] = fmt.Sprint(c.Eval())
	}
	return out
}
