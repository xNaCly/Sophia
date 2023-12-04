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

func indexHelper(parent *Ident, target any, index []Node) any {
	switch v := target.(type) {
	case []interface{}:
		{
			// eg: [array.0]
			in := index[0]
			idx, ok := in.Eval().(float64)
			if !ok {
				t := in.GetToken()
				serror.Add(t, "Index error", "Can't index array with %q, use a number", token.TOKEN_NAME_MAP[t.Type])
				serror.Panic()
			}
			curTarget := v[int(idx)]

			if len(index) == 1 {
				return curTarget
			}

			// eg: [array.0.x]
			return indexHelper(parent, curTarget, index[1:])
		}
	case map[string]interface{}:
		{
			// eq: [map.x]
			in := index[0]
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
			return indexHelper(parent, curTarget, index[1:])
		}
	case nil:
		serror.Add(parent.Token, "Index error", "Can not access nothing (nil)")
		serror.Panic()
	default:
		serror.Add(parent.Token, "Index error", "Element to index into of unknown type %T, not yet implemented", target)
		serror.Panic()
	}
	return nil
}

func (i *Index) Eval() any {
	ident := castPanicIfNotType[*Ident](i.Target, i.Target.GetToken())
	requested, found := consts.SYMBOL_TABLE[ident.Name]
	if !found {
		serror.Add(ident.Token, "Index error", "Requested element %q not defined", ident.Name)
		serror.Panic()
	}
	return indexHelper(ident, requested, i.Index)
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
