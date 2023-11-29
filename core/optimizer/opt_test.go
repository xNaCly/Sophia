package optimizer

import (
	"sophia/core/lexer"
	"sophia/core/parser"
	"testing"
)

func TestOptimizer(t *testing.T) {
	tests := []string{
		// "(fun square (_n) (*n n))",
		// "(let b 5)",
		// "(let b 5)(let b 5)",
		// "(let r 5)(let r 12)",
		// "(fun square (_n) (*n n))(fun square (_n))(fun square (_n))",
		// "(if true)",
		// "(match)",
		// "(for (_ i) 20)",
		// "(fun dummy (_))(put (dummy))",
		// "(fun dummy (_))(let b (dummy))(put b)",
		// "(let b 12)(fun dummy (_))(let b (dummy))(put b)",
	}
	for _, test := range tests {
		tokens := lexer.New(test).Lex()
		ast := parser.New(tokens, "test").Parse()
		ast = New().Start(ast)
		if len(ast) != 0 {
			t.Errorf("Expected ast to be empty, got %#v", ast)
		}
	}
}
