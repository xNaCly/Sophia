package expr

import (
	"fmt"
	"sophia/core/consts"
	"sophia/core/token"
	"strings"
)

// defining a variable
type Var struct {
	Token token.Token
	Ident Node
	Value []Node
}

func (v *Var) GetToken() token.Token {
	return v.Token
}

func (v *Var) Eval() any {
	var val any
	if !(v.Ident.GetToken().Type == token.IDENT || v.Ident.GetToken().Type == token.LEFT_BRACKET) {
		panic("expected an identifier or an array / object index for variable definition, got something else")
	} else if res, ok := v.Ident.(*Ident); ok {
		if len(v.Value) > 1 {
			val = make([]any, len(v.Value))
			for i, c := range v.Value {
				val.([]any)[i] = c.Eval()
			}
		} else if len(v.Value) == 0 {
			val = nil
		} else {
			val = v.Value[0].Eval()
		}
		consts.SYMBOL_TABLE[res.Name] = val
	} else {
		index := castPanicIfNotType[*Index](v.Ident, token.LET)
		ident := castPanicIfNotType[*Ident](index.Element, token.LEFT_BRACKET)
		requested, found := consts.SYMBOL_TABLE[ident.Name]
		if !found {
			panic(fmt.Sprintf("requested element %q not defined", ident.Name))
		}
		switch requested.(type) {
		case []interface{}:
			{
				arr := requested.([]interface{})
				in, ok := index.Index.Eval().(float64)
				if !ok {
					panic(fmt.Sprintf("can't index into array with %T, use a number", index.Index))
				}
				// TODO
				arr[int(in)] = v.Value[0].Eval()
				consts.SYMBOL_TABLE[ident.Name] = arr
			}
		case map[string]interface{}:
			{
				m := requested.(map[string]interface{})
				in, ok := index.Index.(*Ident)
				if !ok {
					panic(fmt.Sprintf("can't index object with %T, use an identifier", index.Index))
				}

				m[in.Name] = v.Value[0].Eval()
				consts.SYMBOL_TABLE[ident.Name] = m
			}
		default:
			panic(fmt.Sprintf("Element to index into of unknown type %T, not yet implemented", requested))
		}
	}

	return val
}

// TODO
func (n *Var) CompileJs(b *strings.Builder) {
	// b.WriteString("let ")
	// b.WriteString(n.Name)
	// if len(n.Value) > 1 {
	// b.WriteString("=")
	// b.WriteRune('[')
	// for i, c := range n.Value {
	// c.CompileJs(b)
	// if i+1 < len(n.Value) {
	// b.WriteRune(',')
	// }
	// }
	// b.WriteRune(']')
	// } else if len(n.Value) == 1 {
	// b.WriteString("=")
	// n.Value[0].CompileJs(b)
	// }
}
