package builtin

import (
	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
)

func builtinAssert(tok *token.Token, args ...types.Node) any {
	if len(args) < 1 {
		serror.Add(tok, "Argument error", "Expected 1 argument for assert builtin")
		serror.Panic()
	}
	if len(args) > 1 {
		serror.Add(args[1].GetToken(), "Argument error", "Too many arguments, expected 1 argument for assert builtin")
		serror.Panic()
	}
	execution := args[0].Eval()
	if res, ok := execution.(bool); !ok {
		serror.Add(args[0].GetToken(), "Type error", "Expected assertion to be of type boolean, got %T", execution)
		serror.Panic()
	} else if !res {
		serror.Add(args[0].GetToken(), "Assertion error", "Assertion failed, wanted true, got false")
		serror.Panic()
	}
	return nil
}
