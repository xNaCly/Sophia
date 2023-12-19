package builtin

import (
	"sophia/core/serror"
	"sophia/core/token"
	"sophia/core/types"
)

func builtinLen(tok *token.Token, args ...types.Node) any {
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
