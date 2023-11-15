package expr

import (
	"sophia/core/consts"
	"sophia/core/debug"
	"sophia/core/serror"
	"sophia/core/token"
	"strings"
)

// function definition
type For struct {
	Token    token.Token
	Params   Node
	LoopOver Node
	Body     []Node
}

func (f *For) GetToken() token.Token {
	return f.Token
}

func (f *For) Eval() any {
	params := f.Params.(*Params).Children
	if len(params) < 1 {
		serror.Add(&f.Token, "Not enough arguments", "Expected at least %d parameters for loop, got %d.", 1, len(params))
		serror.Panic()
	}
	element := castPanicIfNotType[*Ident](params[0], params[0].GetToken())
	oldValue, foundOldValue := consts.SYMBOL_TABLE[element.Name]

	v := f.LoopOver.Eval()
	switch v.(type) {
	case []interface{}:
		loopOver := castPanicIfNotType[[]interface{}](v, f.LoopOver.GetToken())

		for _, el := range loopOver {
			consts.SYMBOL_TABLE[element.Name] = el
			for _, stmt := range f.Body {
				stmt.Eval()
			}
		}
	case float64:
		con := int(v.(float64))
		for i := 0; i < con; i++ {
			consts.SYMBOL_TABLE[element.Name] = i
			for _, stmt := range f.Body {
				stmt.Eval()
			}
		}
	default:
		t := f.LoopOver.GetToken()
		serror.Add(&t, "Invalid iterator", "expected container or upper bound for iteration, got: %T\n", v)
		serror.Panic()
	}

	if foundOldValue {
		consts.SYMBOL_TABLE[element.Name] = oldValue
	}
	return nil
}
func (n *For) CompileJs(b *strings.Builder) {
	// a for loop without any content can be savely removed from execution
	if len(n.Body) == 0 {
		debug.Log("opt: removed 'for loop' with no body at line", n.Token.Line)
		return
	}
	b.WriteString("for(")
	b.WriteString("let ")
	if n.LoopOver.GetToken().Type == token.FLOAT {
		b.WriteString("i = 0; i < ")
		b.WriteString(n.LoopOver.GetToken().Raw)
		b.WriteString("; i++")
	} else {
		// TODO: check if container here
		n.Params.CompileJs(b)
		b.WriteString(" of ")
		n.LoopOver.CompileJs(b)
	}
	b.WriteRune(')')
	b.WriteString("{")
	for _, c := range n.Body {
		c.CompileJs(b)
		b.WriteRune(';')
	}
	b.WriteString("}")
}
