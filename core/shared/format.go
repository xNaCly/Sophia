package shared

import (
	"fmt"
	"github.com/xnacly/sophia/core/types"
	"strconv"
	"strings"
)

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
