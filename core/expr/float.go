package expr

import (
	"sophia/core/serror"
	"sophia/core/token"
	"strconv"
	"strings"
)

type Float struct {
	Token token.Token
}

func (f *Float) GetToken() token.Token {
	return f.Token
}

func (f *Float) Eval() any {
	before, after, found := strings.Cut(f.Token.Raw, "..")
	if found {
		var first float64
		if len(before) == 0 {
			first = 0
		} else {
			var err error
			first, err = strconv.ParseFloat(before, 64)
			if err != nil {
				serror.Add(&f.Token, "Float parse error", "Failed to parse float %q: %q", f.Token.Raw, err)
				serror.Panic()
			}
		}

		if len(after) == 0 {
			serror.Add(&f.Token, "Array spread error", "Upper array spread limit required")
			serror.Panic()
		}

		last, err := strconv.ParseFloat(after, 64)
		if err != nil {
			serror.Add(&f.Token, "Float parse error", "Failed to parse float %q: %q", f.Token.Raw, err)
			serror.Panic()
		}
		r := make([]interface{}, 0)
		for i := first; i < last+1; i++ {
			r = append(r, i)
		}
		return r
	}
	float, err := strconv.ParseFloat(f.Token.Raw, 64)
	if err != nil {
		serror.Add(&f.Token, "Float parse error", "Failed to parse float %q: %q", f.Token.Raw, err)
		serror.Panic()
	}
	return float
}
func (n *Float) CompileJs(b *strings.Builder) {
	b.WriteString(n.Token.Raw)
}
