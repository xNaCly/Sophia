package expr

import (
	"fmt"
	"sophia/core/consts"
	"sophia/core/token"
	"strings"
)

type Index struct {
	Token   token.Token
	Element Node
	Index   Node
}

func (i *Index) GetToken() token.Token {
	return i.Token
}

func (i *Index) Eval() any {
	ident := castPanicIfNotType[*Ident](i.Element, token.LEFT_BRACKET)
	requested, found := consts.SYMBOL_TABLE[ident.Name]
	if !found {
		panic(fmt.Sprintf("requested element %q not defined", ident.Name))
	}
	switch requested.(type) {
	case []interface{}:
		{
			arr := requested.([]interface{})
			index, ok := i.Index.Eval().(float64)
			if !ok {
				panic(fmt.Sprintf("can't index array with %T, use a number", i.Index))
			}
			return arr[int(index)]
		}
	case map[string]interface{}:
		{
			m := requested.(map[string]interface{})
			index, ok := i.Index.(*Ident)
			if !ok {
				panic(fmt.Sprintf("can't index object with %T, use an identifier", i.Index))
			}
			return m[index.Name]
		}
	default:
		panic(fmt.Sprintf("Element to index into of unknown type %T, not yet implemented", requested))
	}
}

// TODO:
func (i *Index) CompileJs(b *strings.Builder) {}
