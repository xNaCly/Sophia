package core

var SYMBOL_TABLE = map[string]any{}

func Eval(ast []Node) []float64 {
	out := make([]float64, 0)
	for _, c := range ast {
		if val, ok := c.Eval().(float64); ok {
			out = append(out, val)
		}
	}
	return out
}
