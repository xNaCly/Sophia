package expr

import (
	"sophia/core/serror"
	"sophia/core/token"
)

// attempts to cast `in` to `T`, returns `in` cast to `T` if successful. If
// cast fails, panics.
func castPanicIfNotType[T any](in any, t token.Token) T {
	val, ok := in.(T)
	if !ok {
		var e T
		serror.Add(&t, "Type error", "Incompatiable types %T and %T", in, e)
		serror.Panic()
	}
	return val
}
