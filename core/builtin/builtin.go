// builtin provides functions that are built into the sophia language but are
// written in pure go, they may interface with the sophia lang via AST
// manipulation and by accepting AST nodes and returning values or nodes.
//
// See docs/Embedding.md for more information.
package builtin

import (
	"os"
	"sophia/core/alloc"
	"sophia/core/consts"
	"sophia/core/expr"
	"sophia/core/serror"
	"sophia/core/shared"
	"sophia/core/token"
	"sophia/core/types"
	"strings"
)

var sharedPrintBuffer = &strings.Builder{}

func init() {
	consts.FUNC_TABLE[alloc.NewFunc("len")] = func(tok *token.Token, args ...types.Node) any {
		if len(args) < 1 || len(args) > 1 {
			serror.Add(tok, "Argument error", "Expected at least and at most 1 argument for len built-in")
			serror.Panic()
		}
		// the compiler is somehow not smart enough to let me write string, []any, etc...
		switch v := args[0].Eval().(type) {
		case string:
			return len(v)
		case map[string]any:
			return len(v)
		case []any:
			return len(v)
		default:
			serror.Add(tok, "Error", "Can't compute length for target of type %T", v)
			serror.Panic()
		}
		return nil
	}
	consts.FUNC_TABLE[alloc.NewFunc("println")] = func(tok *token.Token, args ...types.Node) any {
		sharedPrintBuffer.Reset()
		shared.FormatHelper(sharedPrintBuffer, args, ' ')
		sharedPrintBuffer.WriteRune('\n')
		os.Stdout.WriteString(sharedPrintBuffer.String())
		return nil
	}
	consts.FUNC_TABLE[alloc.NewFunc("map")] = func(tok *token.Token, args ...types.Node) any {
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

		if len(call.Params) != 0 {
			serror.Add(args[0].GetToken(), "Argument Error", "Expected function call with 0 arguments, got %d", len(call.Params))
			serror.Panic()
		}

		call.Params = make([]types.Node, 1)

		var r any
		switch iter := args[1].Eval().(type) {
		// string requires a copy, sadly
		case string:
			t := make([]rune, len(iter))
			for i, char := range iter {
				call.Params[0] = &expr.Float{Value: float64(char)}
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
			for i, element := range iter {
				call.Params[0] = &expr.Any{Value: element}
				iter[i] = call.Eval()
			}
			return iter
		default:
			serror.Add(args[1].GetToken(), "Error", "Can't map over target of type %T, expected string, array or object", args[1])
			serror.Panic()
		}

		return r
	}
}
