package builtin

import (
	"github.com/xnacly/sophia/core/expr"
	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
)

func builtinMap(tok *token.Token, args ...types.Node) any {
	if len(args) != 2 {
		serror.Add(tok, "Argument error", "Expected exactly 2 arguments for map built-in, first function, second iterator")
		serror.Panic()
	}

	// function to apply to iterator
	switch args[0].(type) {
	case *expr.Call, *expr.Lambda:
	default:
		serror.Add(args[0].GetToken(), "Argument Error", "Expected first argument to be a function call, got %T", args[0])
		serror.Panic()
	}

	call := args[0]

	var r any
	switch iter := args[1].Eval().(type) {
	// string requires a copy, sadly
	case string:
		t := make([]float64, len(iter))
		for i, char := range iter {
			call.SetChildren([]types.Node{&expr.Float{Value: float64(char)}})
			res := call.Eval()
			out, ok := res.(float64)
			if !ok {
				serror.Add(call.GetToken(), "Type error", "Expected result of type float64 for function used for string mapping, got %T instead", res)
				serror.Panic()
			}
			t[i] = out
		}
		r = t
	case []any:
		t := make([]any, len(iter))
		for i, element := range iter {
			call.SetChildren([]types.Node{&expr.Any{Value: element}})
			t[i] = call.Eval()
		}
		r = t
	default:
		serror.Add(args[1].GetToken(), "Error", "Can't map over target of type %T, expected string, array or object", args[1])
		serror.Panic()
	}

	return r
}
