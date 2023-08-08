package expr

import (
	"fmt"
	"sophia/core/token"
)

// attempts to cast `in` to `T`, returns `in` cast to `T` if successful. If
// cast fails, panics.
func castPanicIfNotType[T any](in any, op int) T {
	val, ok := isType[T](in)
	if !ok {
		panic(fmt.Sprintf("can not use variable of type %T in current operation (%s), expected %T for value %+v", in, token.TOKEN_NAME_MAP[op], val, in))
	}
	return val
}

// checks if `in` is castable to `T`, returns casted value and true if
// castable, zero value of `T` and false if not
func isType[T any](in any) (T, bool) {
	val, ok := in.(T)
	if !ok {
		var e T
		return e, false
	}
	return val, true
}

func extractChild(n Node, op int) float64 {
	var val float64
	if idt, ok := isType[*Ident](n); ok {
		arr := castPanicIfNotType[[]interface{}](idt.Eval(), op)
		for i, item := range arr {
			t := castPanicIfNotType[float64](item, op)
			if i == 0 {
				val = t
				continue
			}
			switch op {
			case token.ADD:
				val += t
			case token.SUB:
				val -= t
			case token.DIV:
				val /= t
			case token.MUL:
				val *= t
			case token.MOD:
				val = float64(int(val) % int(t))
			}
		}
	} else {
		val = castPanicIfNotType[float64](n.Eval(), op)
	}
	return val
}
