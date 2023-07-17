package core

func Eval(ast []Node) []float64 {
	out := make([]float64, 0)
	for _, c := range ast {
		out = append(out, c.Eval())
	}
	return out
}
