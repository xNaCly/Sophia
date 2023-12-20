package builtin

import (
	"sophia/core/expr"
	"sophia/core/serror"
	"sophia/core/token"
	"sophia/core/types"
)

func buildinMap(tok *token.Token, args ...types.Node) any {
	if len(args) != 2 {
		serror.Add(tok, "Argument error", "Expected exactly 2 arguments for map built-in")
		serror.Panic()
	}

	// function to apply to iterator
	call, ok := args[0].(*expr.Call)
	if !ok {
		serror.Add(args[0].GetToken(), "Argument Error", "Expected first argument to be a function call, got %T", args[0])
		serror.Panic()
	}

	if len(call.Args) != 0 {
		serror.Add(args[0].GetToken(), "Argument Error", "Expected function call with 0 arguments, got %d", len(call.Args))
		serror.Panic()
	}

	call.Args = make([]types.Node, 1)

	var r any
	switch iter := args[1].Eval().(type) {
	// string requires a copy, sadly
	case string:
		t := make([]rune, len(iter))
		for i, char := range iter {
			call.Args[0] = &expr.Float{Value: float64(char)}
			res := call.Eval()
			out, ok := res.(float64)
			if !ok {
				serror.Add(call.Token, "Type error", "Expected result of type float64 for function used for string mapping, got %T instead", res)
				serror.Panic()
			}
			t[i] = rune(out)
		}
		r = string(t)
	case []any:
		t := make([]any, len(iter))
		for i, element := range iter {
			call.Args[0] = &expr.Any{Value: element}
			t[i] = call.Eval()
		}
		r = t
	default:
		serror.Add(args[1].GetToken(), "Error", "Can't map over target of type %T, expected string, array or object", args[1])
		serror.Panic()
	}

	return r
}
