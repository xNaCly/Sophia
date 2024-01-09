package builtin

import (
	"github.com/xnacly/sophia/core/expr"
	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
)

func builtinFilter(tok *token.Token, args ...types.Node) any {
	if len(args) != 2 {
		serror.Add(tok, "Argument error", "Expected exactly 2 arguments for filter built-in, first function and second iterator")
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
		t := make([]rune, 0, len(iter))
		for _, char := range iter {
			call.SetChildren([]types.Node{&expr.Float{Value: float64(char)}})
			res := call.Eval()
			out, ok := res.(bool)
			if !ok {
				serror.Add(call.GetToken(), "Type error", "Expected result of type bool for function used for filter, got %T instead", res)
				serror.Panic()
			}
			if out {
				t = append(t, char)
			}
		}
		r = string(t)
	case []any:
		t := make([]any, 0, len(iter))
		for _, element := range iter {
			call.SetChildren([]types.Node{&expr.Any{Value: element}})
			res := call.Eval()
			out, ok := res.(bool)
			if !ok {
				serror.Add(call.GetToken(), "Type error", "Expected result of type bool for function used for filter, got %T instead", res)
				serror.Panic()
			}
			if out {
				t = append(t, element)
			}
		}
		r = t
	default:
		serror.Add(args[1].GetToken(), "Error", "Can't filter target of type %T, expected string or array", args[1])
		serror.Panic()
	}

	return r
}
