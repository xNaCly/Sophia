package consts

import (
	"fmt"
	"os"
	"sophia/core/alloc"
	"sophia/core/serror"
	"sophia/core/token"
	"sophia/core/types"
	"strconv"
	"strings"
)

type Return struct {
	HasValue bool
	Value    any
}

var RETURN = Return{}

var sharedPrintBuffer = &strings.Builder{}

var FUNC_TABLE = map[uint32]any{
	alloc.NewFunc("len"): func(token *token.Token, n ...types.Node) any {
		if len(n) < 1 || len(n) > 1 {
			serror.Add(token, "Argument error", "Expected at least and at most 1 argument for len built-in")
			serror.Panic()
		}
		// the compiler is somehow not smart enough to let me write string, []any, etc...
		switch v := n[0].Eval().(type) {
		case string:
			return len(v)
		case map[string]any:
			return len(v)
		case []any:
			return len(v)
		default:
			serror.Add(token, "Error", "Can't compute length for target of type %T", v)
			serror.Panic()
		}
		return nil
	},
	alloc.NewFunc("println"): func(token *token.Token, n ...types.Node) any {
		sharedPrintBuffer.Reset()
		FormatHelper(sharedPrintBuffer, n, ' ')
		sharedPrintBuffer.WriteRune('\n')
		os.Stdout.WriteString(sharedPrintBuffer.String())
		return nil
	},
}

// formats the given children by executing them, skips fmt.Sprint for string,
// float64 and booleans. Uses a passed in buffer for skipping memory
// allocation for each call. Remember to reset the buffer before calling this
// function.
func FormatHelper(buffer *strings.Builder, children []types.Node, sep rune) {
	for i, c := range children {
		if i != 0 && sep != 0 {
			buffer.WriteRune(sep)
		}
		v := c.Eval()
		switch v := v.(type) {
		case string:
			buffer.WriteString(v)
		case float64:
			buffer.WriteString(strconv.FormatFloat(v, 'g', 12, 64))
		case bool:
			if v {
				buffer.WriteString("true")
			} else {
				buffer.WriteString("false")
			}
		default:
			fmt.Fprint(buffer, v)
		}
	}
}
