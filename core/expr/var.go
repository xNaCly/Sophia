package expr

import (
	"sophia/core/consts"
	"sophia/core/serror"
	"sophia/core/token"
	"strings"
)

// defining a variable
type Var struct {
	Token *token.Token
	Ident Node
	Value []Node
}

func (v *Var) GetToken() *token.Token {
	return v.Token
}

func (v *Var) Eval() any {
	var val any
	if !(v.Ident.GetToken().Type == token.IDENT || v.Ident.GetToken().Type == token.LEFT_BRACKET) {
		t := v.Ident.GetToken()
		serror.Add(t, "Variable error", "Expected an identifier, an array or object index for variable definition, got something else")
		serror.Panic()
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
		index := castPanicIfNotType[*Index](v.Ident, v.Ident.GetToken())
		ident := castPanicIfNotType[*Ident](index.Element, index.Element.GetToken())
		requested, found := consts.SYMBOL_TABLE[ident.Name]
		if !found {
			t := v.Ident.GetToken()
			serror.Add(t, "Undefined variable", "Requested item %q not found", ident.Name)
			serror.Panic()
		}
		switch requested.(type) {
		case []interface{}:
			{
				arr := requested.([]interface{})
				in, ok := index.Index.Eval().(float64)
				if !ok {
					t := index.Index.GetToken()
					serror.Add(t, "Index error", "Can't index array with %q, use a number", token.TOKEN_NAME_MAP[t.Type])
					serror.Panic()
				}
				arr[int(in)] = v.Value[0].Eval()
				consts.SYMBOL_TABLE[ident.Name] = arr
			}
		case map[string]interface{}:
			{
				m := requested.(map[string]interface{})
				in, ok := index.Index.(*Ident)
				if !ok {
					t := index.GetToken()
					serror.Add(t, "Index error", "Can't index object with %q, use an identifier", token.TOKEN_NAME_MAP[t.Type])
					serror.Panic()
				}

				m[in.Name] = v.Value[0].Eval()
				consts.SYMBOL_TABLE[ident.Name] = m
			}
		default:
			serror.Add(ident.Token, "Index error", "Element to index into of unknown type %T, not yet implemented", requested)
			serror.Panic()
		}
	}

	return val
}

func (n *Var) CompileJs(b *strings.Builder) {
	ident, _ := n.Ident.(*Ident)
	if ident == nil {
		return
	}
	// js does not want let for already declared variables
	if _, ok := consts.SYMBOL_TABLE[ident.Name]; !ok {
		consts.SYMBOL_TABLE[ident.Name] = true
		b.WriteString("let ")
	}
	ident.CompileJs(b)
	if len(n.Value) > 1 {
		b.WriteString("=")
		b.WriteRune('[')
		for i, c := range n.Value {
			if c == nil {
				continue
			}
			c.CompileJs(b)
			if i+1 < len(n.Value) {
				b.WriteRune(',')
			}
		}
		b.WriteRune(']')
	} else if len(n.Value) == 1 {
		v := n.Value[0]
		if v == nil {
			return
		}
		b.WriteString("=")
		v.CompileJs(b)
	}
}
