package expr

import (
	"bytes"
	"fmt"
	"sophia/core/serror"
	"sophia/core/token"
	"strconv"
)

// formats the given children by executing them, skips fmt.Sprint for string,
// float64 and booleans. Uses a passed in buffer for skipping memory
// allocation for each call. Remember to reset the buffer before calling this
// function.
func formatHelper(buffer *bytes.Buffer, children []Node) {
	for i, c := range children {
		if i != 0 {
			buffer.WriteRune(' ')
		}
		v := c.Eval()
		switch v := v.(type) {
		case string:
			buffer.WriteString(v)
		case float64:
			buffer.WriteString(strconv.FormatFloat(v, 'g', 2, 64))
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

// fastpath for casting bool, reduces memory allocation by skipping allocation
func castBoolPanic(in any, t *token.Token) bool {
	switch v := in.(type) {
	case bool:
		return v
	default:
		serror.Add(t, "Type error", "Incompatiable types %T and bool", in)
		serror.Panic()
	}
	// technically unreachable
	return false
}

// fastpath for casting float64, reduces memory allocation by skipping allocation
func castFloatPanic(in any, t *token.Token) float64 {
	switch v := in.(type) {
	case float64:
		return v
	default:
		serror.Add(t, "Type error", "Incompatiable types %T and float64", in)
		serror.Panic()
	}
	// technically unreachable
	return 0
}

// attempts to cast `in` to `T`, returns `in` cast to `T` if successful. If
// cast fails, panics.
func castPanicIfNotType[T any](in any, t *token.Token) T {
	val, ok := in.(T)
	if !ok {
		var e T
		serror.Add(t, "Type error", "Incompatiable types %T and %T", in, e)
		serror.Panic()
	}
	return val
}
