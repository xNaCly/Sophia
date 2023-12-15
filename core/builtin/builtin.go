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
}
