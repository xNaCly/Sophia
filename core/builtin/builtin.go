// builtin provides functions that are built into the sophia language but are
// written in pure go, they may interface with the sophia lang via AST
// manipulation and by accepting AST nodes and returning values or nodes.
//
// See docs/Embedding.md for more information.
package builtin

import (
	"sophia/core/alloc"
	"sophia/core/consts"
	"sophia/core/types"
)

func init() {
	builtins := map[string]types.KnownFunctionInterface{
		"len":     builtinLen,
		"println": buildinPrintln,
		"map":     buildinMap,
	}

	for name, function := range builtins {
		consts.FUNC_TABLE[alloc.NewFunc(name)] = function
	}
}
