package builtin

import (
	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
)

func builtinType(tok *token.Token, args ...types.Node) any {
	if len(args) < 1 {
		serror.Add(tok, "Argument error", "Expected 1 argument for assert builtin")
		serror.Panic()
	}
	if len(args) > 1 {
		serror.Add(args[1].GetToken(), "Argument error", "Too many arguments, expected 1 argument for assert builtin")
		serror.Panic()
	}

	// TODO: add all missing types
	switch args[0].Eval().(type) {
	case []any:
		return "array"
	case map[string]any:
		return "object"
	case float64:
		return "float"
	case string:
		return "string"
	default:
		serror.Add(args[0].GetToken(), "Not implemented", "type built-in Not implemented for %T", args[0])
		serror.Panic()
		return nil
	}
}
