package expr

import (
	"sophia/core/consts"
	"sophia/core/serror"
	"sophia/core/token"
	"strings"
)

type Index struct {
	Token  *token.Token
	Target Node
	Index  []Node
}

func (i *Index) GetChildren() []Node {
	return nil
}

func (n *Index) SetChildren(c []Node) {}
func (i *Index) GetToken() *token.Token {
	return i.Token
}

func indexHelper(target any, index []Node) any {
	switch v := target.(type) {
	case []interface{}:
		{
			// eg: [array.0]
			in := index[0]
			switch V := in.(type) {
			case *Ident:
				t := in.GetToken()
				serror.Add(t, "Index error", "Can't index array.%s, not an object", V.Name)
				serror.Panic()
			}
			idxf, ok := in.Eval().(float64)
			if !ok {
				t := in.GetToken()
				serror.Add(t, "Index error", "Can't index array with %q, use a number", token.TOKEN_NAME_MAP[t.Type])
				serror.Panic()
			}
			idx := int(idxf)
			if idx > len(v) {
				serror.Add(in.GetToken(), "Out of bounds error", "Array has length of %d, index %d can not be accessed", len(v), idx)
				serror.Panic()
			}
			curTarget := v[int(idx)]

			if len(index) == 1 {
				return curTarget
			}

			// eg: [array.0.x]
			return indexHelper(curTarget, index[1:])
		}
	case map[string]interface{}:
		{
			// eq: [map.x]
			in := index[0]
			switch V := in.(type) {
			case *Float:
				t := in.GetToken()
				serror.Add(t, "Index error", "Can't index object.%g, not an array", V.Value)
				serror.Panic()
			}
			idx, ok := in.(*Ident)
			if !ok {
				t := idx.GetToken()
				serror.Add(t, "Index error", "Can't index object with %q, use a string", token.TOKEN_NAME_MAP[t.Type])
				serror.Panic()
			}
			curTarget := v[idx.Name]
			if len(index) == 1 {
				return curTarget
			}

			// eg: [map.x.y]
			return indexHelper(curTarget, index[1:])
		}
	case nil:
		// TODO: display what part of the index is nil: person.bank.etc
		//                                                     ^^^^ is null, thus .etc will error
		val := index[0].Eval()
		serror.Add(index[0].GetToken(), "Index error", "Index %v unavailable on %v", val, target)
		serror.Panic()
	default:
		switch V := index[0].(type) {
		case *Ident:
			serror.Add(index[0].GetToken(), "Index error", "Not an object, can't use <target>.%s", V.Name)
			serror.Panic()
		case *Float:
			serror.Add(index[0].GetToken(), "Index error", "Not an array, can't use <target>.%g", V.Value)
			serror.Panic()
		}
	}
	return nil
}

func (i *Index) Eval() any {
	ident := castPanicIfNotType[*Ident](i.Target, i.Target.GetToken())
	requested, found := consts.SYMBOL_TABLE[ident.Key]
	if !found {
		serror.Add(ident.Token, "Index error", "Requested element %q not defined", ident.Name)
		serror.Panic()
	}
	return indexHelper(requested, i.Index)
}

func (i *Index) CompileJs(b *strings.Builder) {
	i.Target.CompileJs(b)
	for _, index := range i.Index {
		b.WriteRune('[')
		switch v := index.(type) {
		case *Ident:
			b.WriteRune('"')
			v.CompileJs(b)
			b.WriteRune('"')
		case *Float:
			v.CompileJs(b)
		default:
			t := v.GetToken()
			serror.Add(t, "Index error", "Element to index into of unknown type, not yet implemented")
			serror.Panic()
		}
		b.WriteRune(']')
	}
}
