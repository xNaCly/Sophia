package expr

import (
	"fmt"
	"sophia/core/consts"
	"sophia/core/token"
	"strings"
)

// function definition
type Func struct {
	Token  token.Token
	Name   Node
	Params Node
	Body   []Node
}

func (f *Func) GetToken() token.Token {
	return f.Token
}

func (f *Func) Eval() any {
	if f.Name.GetToken().Type == token.IDENT {
		consts.FUNC_TABLE[f.Name.GetToken().Raw] = f
		return nil
	} else if f.Name.GetToken().Type == token.LEFT_BRACKET {
		i, _ := f.Name.(*Index)
		ident := castPanicIfNotType[*Ident](i.Element, token.FUNC)
		index := castPanicIfNotType[*Ident](i.Index, token.FUNC)
		requested, found := consts.SYMBOL_TABLE[ident.Name]
		if !found {
			panic(fmt.Sprintf("can't define function %q on not existing object %q", index.Name, ident.Name))
		}
		requestedObject := castPanicIfNotType[map[string]interface{}](requested, token.FUNC)
		requestedObject[index.Name] = f
	}
	return nil
}

func (n *Func) CompileJs(b *strings.Builder) {
	cLen := len(n.Body)
	b.WriteString("function ")
	b.WriteString(n.Name.GetToken().Raw)
	b.WriteRune('(')
	n.Params.CompileJs(b)
	b.WriteString("){")
	for i, c := range n.Body {
		if i+1 == cLen {
			b.WriteString("return ")
		}
		c.CompileJs(b)
		b.WriteRune(';')
	}
	b.WriteRune('}')
}
