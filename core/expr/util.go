package expr

import (
	"fmt"
	"sophia/core/token"
)

// attempts to cast `in` to `T`, returns `in` cast to `T` if successful. If
// cast fails, panics.
func castPanicIfNotType[T any](in any, op int) T {
	val, ok := in.(T)
	if !ok {
		panic(fmt.Sprintf("can not use variable of type %T in current operation (%s), expected %T for value %+v", in, token.TOKEN_NAME_MAP[op], val, in))
	}
	return val
}
