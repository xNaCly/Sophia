package expr

import (
	"sophia/core/consts"
	"sophia/core/serror"
	"sophia/core/token"
	"strings"
)

type Index struct {
	Token   *token.Token
	Element Node
	Index   Node
}

func (i *Index) GetToken() *token.Token {
	return i.Token
}

func (i *Index) Eval() any {
	ident := castPanicIfNotType[*Ident](i.Element, i.Element.GetToken())
	requested, found := consts.SYMBOL_TABLE[ident.Name]
	if !found {
		serror.Add(ident.Token, "Index error", "Requested element %q not defined", ident.Name)
		serror.Panic()
	}

	switch requested.(type) {
	case []interface{}:
		{
			arr := requested.([]interface{})
			index, ok := i.Index.Eval().(float64)
			if !ok {
				t := i.Index.GetToken()
				serror.Add(t, "Index error", "Can't index array with %q, use a number", token.TOKEN_NAME_MAP[t.Type])
				serror.Panic()
			}
			return arr[int(index)]
		}
	case map[string]interface{}:
		{
			m := requested.(map[string]interface{})
			index, ok := i.Index.(*Ident)
			if !ok {
				t := i.Index.GetToken()
				serror.Add(t, "Index error", "Can't index object with %q, use a number", token.TOKEN_NAME_MAP[t.Type])
				serror.Panic()
			}
			return m[index.Name]
		}
	case nil:
		serror.Add(ident.Token, "Index error", "Can not access nothing (nil)")
		serror.Panic()
	default:
		serror.Add(ident.Token, "Index error", "Element to index into of unknown type %T, not yet implemented", requested)
		serror.Panic()
	}
	return nil
}

func (i *Index) CompileJs(b *strings.Builder) {
	i.Element.CompileJs(b)
	b.WriteRune('[')
	switch i.Index.(type) {
	case *Ident:
		b.WriteRune('"')
		i.Index.CompileJs(b)
		b.WriteRune('"')
	case *Float:
		i.Index.CompileJs(b)
	default:
		t := i.Index.GetToken()
		serror.Add(t, "Index error", "Element to index into of unknown type, not yet implemented")
		serror.Panic()
	}
	b.WriteRune(']')
}
