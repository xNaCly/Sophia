package builtin

import (
	"github.com/xnacly/sophia/core/expr"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
	"testing"
)

// TODO: more tests

func TestBuiltInLen(t *testing.T) {
	tests := []struct {
		name  string
		len   int
		input types.Node
	}{
		{name: "string", input: &expr.String{Token: &token.Token{Raw: "1234"}}, len: 4},
		{name: "empty array", input: &expr.Array{Children: []types.Node{}}, len: 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := builtinLen(nil, test.input)
			if test.len != r {
				t.Errorf("Expected %d, got %d for %#v", test.len, r, test.input)
			}
		})
	}
}
