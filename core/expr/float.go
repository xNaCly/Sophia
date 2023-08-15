package expr

import (
	"fmt"
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
	// TODO: maybe move this to the parser
	before, after, found := strings.Cut(f.Token.Raw, "..")
	if found {
		var first float64
		if len(before) == 0 {
			first = 0
		} else {
			var err error
			first, err = strconv.ParseFloat(before, 64)
			if err != nil {
				panic(fmt.Sprint("failed to parse float: ", err))
			}
		}

		if len(after) == 0 {
			panic("upper bound of array spread operator required")
		}

		last, err := strconv.ParseFloat(after, 64)
		if err != nil {
			panic(fmt.Sprint("failed to parse float: ", err))
		}
		r := make([]interface{}, 0)
		for i := first; i < last+1; i++ {
			r = append(r, i)
		}
		return r
	}
	float, err := strconv.ParseFloat(f.Token.Raw, 64)
	if err != nil {
		panic(fmt.Sprint("failed to parse float: ", err))
	}
	return float
}
