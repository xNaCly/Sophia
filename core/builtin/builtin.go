// builtin provides functions that are built into the sophia language but are
// written in pure go, they may interface with the sophia lang via AST
// manipulation and by accepting AST nodes and returning values or nodes.
//
// See docs/Embedding.md for more information.
package builtin

import (
	"github.com/xnacly/sophia/core/alloc"
	"github.com/xnacly/sophia/core/consts"
	"github.com/xnacly/sophia/core/types"
)

func init() {
	builtins := map[string]types.KnownFunctionInterface{
		"len":     builtinLen,
		"map":     builtinMap,
		"type":    builtinType,
		"println": builtinPrintln,
		"filter":  builtinFilter,
		"assert":  builtinAssert,
	}

	for name, function := range builtins {
		consts.FUNC_TABLE[alloc.NewFunc(name)] = function
	}
}
