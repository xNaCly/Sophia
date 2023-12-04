package parser

import (
	"io"
	"os"
	"sophia/core/expr"
	"sophia/core/lexer"
	"sophia/core/serror"
	"sophia/core/token"
	"strconv"
	"strings"
)

type Parser struct {
	token    []*token.Token
	filename string
	pos      int
}

func New(tokens []*token.Token, filename string) *Parser {
	if len(tokens) == 0 {
		serror.Add(&token.Token{LinePos: 0, Raw: " "}, "Unexpected end of input", "Source possibly empty")
		return &Parser{}
	}
	return &Parser{
		token:    tokens,
		pos:      0,
		filename: filename,
	}
}

func (p *Parser) Parse() []expr.Node {
	res := make([]expr.Node, 0)
	for !p.peekIs(token.EOF) {
		stmt := p.parseStatment()
		if stmt != nil && stmt.GetToken().Type == token.LOAD {
			if loadStmt, ok := stmt.(*expr.Load); ok {
				res = append(res, p.loadNewSource(loadStmt)...)
			}
			continue
		}
		if stmt == nil {
			return res
		}
		res = append(res, stmt)
	}
	return res
}

func (p *Parser) loadNewSource(node *expr.Load) []expr.Node {
	res := make([]expr.Node, 0)
	for i := 0; i < len(node.Imports); i++ {
		name := node.Imports[i].GetToken()
		file, err := os.Open(name.Raw)
		if err != nil {
			serror.Add(name, "Failed to source import", "Couldn't open %q: %q.", name.Raw, err)
			continue
		}
		content, err := io.ReadAll(file)
		if err != nil {
			serror.Add(name, "Failed to read import", "Couldn't read %q: %q.", name.Raw, err)
			continue
		}
		lexer := lexer.New(string(content))
		token := lexer.Lex()
		if name.Raw == p.filename {
			serror.Add(name, "Detected recursion in file imports", "Got %q while already parsing %q.", name.Raw, p.filename)
			continue
		}
		parser := New(token, name.Raw)
		res = append(res, parser.Parse()...)
	}
	return res
}

// INFO: @SamuelScheit fixed this, i dont even know how and he doesnt either
func (p *Parser) parseStatment() expr.Node {
	childs := make([]expr.Node, 0)
	var stmt expr.Node
	p.peekError(token.LEFT_BRACE, "Missing statement start")
	p.advance()
	p.peekErrorMany("Missing or unknown operator", token.EXPECTED_KEYWORDS...)
	op := p.peek()
	p.advance()

	for {
		var child expr.Node
		if p.peekIs(token.EOF) || p.peekIs(token.RIGHT_BRACE) {
			break
		} else if p.peekIs(token.LEFT_BRACE) {
			nStmt := p.parseStatment()
			if nStmt == nil {
				return nil
			}
			childs = append(childs, nStmt)
			continue
		} else {
			child = p.parseArguments()
		}

		if child != nil {
			childs = append(childs, child)
		}

		p.advance()
	}

	switch op.Type {
	case token.RETURN:
		var child expr.Node
		if len(childs) > 1 {
			serror.Add(op, "Too many arguments", "Expected zero or one argument to return, got %d.", len(childs))
			return nil
		} else if len(childs) == 1 {
			child = childs[0]
		}
		stmt = &expr.Return{
			Token: op,
			Child: child,
		}
	case token.MATCH:
		stmt = &expr.Match{
			Token:    op,
			Branches: childs,
		}
	case token.LOAD:
		if len(childs) == 0 {
			serror.Add(op, "Not enough arguments", "Expected at least one argument for loading files, got %d.", len(childs))
			return nil
		}
		for _, c := range childs {
			if c.GetToken().Type != token.STRING {
				t := c.GetToken()
				serror.Add(t, "Type error", "Expected an argument of type string for loading files, got %q.", token.TOKEN_NAME_MAP[t.Type])
				return nil
			}
		}
		stmt = &expr.Load{
			Token:   op,
			Imports: childs[0:],
		}
	case token.FOR:
		if len(childs) < 2 {
			serror.Add(op, "Not enough arguments", "Expected two argument for loop definition, got %d.", len(childs))
			return nil
		}
		params := childs[0]
		t := params.GetToken()
		if params.GetToken().Type != token.PARAM {
			serror.Add(t, "Type error", "Expected the first argument for loop definition to be of type PARAM, got %q.", token.TOKEN_NAME_MAP[childs[0].GetToken().Type])
			return nil
		}
		if len(params.(*expr.Params).Children) != 1 {
			serror.Add(t, "Not enough parameters", "Expected one parameter for loop parameter definition, got %d.", len(params.(*expr.Params).Children))
			return nil
		}
		stmt = &expr.For{
			Token:    op,
			Params:   params,
			LoopOver: childs[1],
			Body:     childs[2:],
		}
	case token.IDENT:
		stmt = &expr.Call{
			Token:  op,
			Params: childs,
		}
	case token.LT:
		if len(childs) != 2 {
			serror.Add(op, "Incorrect parameter amount", "Expected exactly two statements for less than comparison, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Lt{
			Token:    op,
			Children: childs,
		}
	case token.GT:
		if len(childs) != 2 {
			serror.Add(op, "Incorrect parameter amount", "Expected exactly two statements for greater than comparison, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Gt{
			Token:    op,
			Children: childs,
		}
	case token.PARAM:
		for _, c := range childs {
			t := c.GetToken()
			if t.Type != token.IDENT {
				serror.Add(t, "Type error", "Expected identifier for parameter definition, got %q.", token.TOKEN_NAME_MAP[t.Type])
				return nil
			}
		}
		stmt = &expr.Params{
			Token:    op,
			Children: childs,
		}
	case token.FUNC:
		if len(childs) < 2 {
			serror.Add(op, "Not enough parameters", "Expected 2 parameters, one for function name and one for parameters, got %d.", len(childs))
			return nil
		}
		t := childs[0].GetToken()
		if t.Type != token.IDENT {
			serror.Add(t, "Type error", "Expected the first argument for function definition to be of type IDENT, got %q.", token.TOKEN_NAME_MAP[t.Type])
			return nil
		}
		t = childs[1].GetToken()
		if childs[1].GetToken().Type != token.PARAM {
			serror.Add(t, "Type error", "Expected the second argument for function definition to be of type PARAM, got %q.", token.TOKEN_NAME_MAP[t.Type])
			return nil
		}
		stmt = &expr.Func{
			Token:  op,
			Name:   childs[0],
			Params: childs[1],
			Body:   childs[2:],
		}
	case token.IF:
		if len(childs) == 0 {
			serror.Add(op, "Not enough arguments", "Expected at least two arguments for condition, got %d.", len(childs))
			return nil
		}
		cond := childs[0]
		stmt = &expr.If{
			Token:     op,
			Condition: cond,
			Body:      childs[1:],
		}
	case token.LET:
		if len(childs) == 0 {
			serror.Add(op, "Not enough arguments", "Expected at least one argument for variable declaration, got %d.", len(childs))
			return nil
		}
		ident := childs[0]
		stmt = &expr.Var{
			Token: op,
			Ident: ident,
			Value: childs[1:],
		}
	case token.MERGE:
		if len(childs) < 2 {
			serror.Add(op, "Incorrect parameter amount", "Expected at least two arguments for merge, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Merge{
			Token:    op,
			Children: childs,
		}
	case token.EQUAL:
		if len(childs) < 2 {
			serror.Add(op, "Incorrect parameter amount", "Expected at least two arguments for equality check, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Equal{
			Token:    op,
			Children: childs,
		}
	case token.NEG:
		if len(childs) != 1 {
			serror.Add(op, "Incorrect parameter amount", "Expected exactly one argument for negation, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Neg{
			Token:    op,
			Children: childs[0],
		}
	case token.OR:
		if len(childs) < 2 {
			serror.Add(op, "Incorrect parameter amount", "Expected 2 or more arguments for or, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Or{
			Token:    op,
			Children: childs,
		}
	case token.AND:
		if len(childs) < 2 {
			serror.Add(op, "Incorrect parameter amount", "Expected 2 or more arguments for and, got %d.", len(childs))
			return nil
		}
		stmt = &expr.And{
			Token:    op,
			Children: childs,
		}
	case token.ADD:
		if len(childs) < 2 {
			serror.Add(op, "Incorrect parameter amount", "Expected 2 or more arguments for addition, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Add{
			Token:    op,
			Children: childs,
		}
	case token.SUB:
		if len(childs) < 2 {
			serror.Add(op, "Incorrect parameter amount", "Expected 2 or more arguments for subtraction, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Sub{
			Token:    op,
			Children: childs,
		}
	case token.DIV:
		if len(childs) < 2 {
			serror.Add(op, "Incorrect parameter amount", "Expected 2 or more arguments for division, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Div{
			Token:    op,
			Children: childs,
		}
	case token.MUL:
		if len(childs) < 2 {
			serror.Add(op, "Incorrect parameter amount", "Expected 2 or more arguments for multiplication, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Mul{
			Token:    op,
			Children: childs,
		}
	case token.MOD:
		if len(childs) < 2 {
			serror.Add(op, "Incorrect parameter amount", "Expected 2 or more arguments for mod, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Mod{
			Token:    op,
			Children: childs,
		}
	case token.PUT:
		if len(childs) == 0 {
			serror.Add(op, "Incorrect parameter amount", "Expected 1 or more arguments for put, got %d.", len(childs))
			return nil
		}
		stmt = &expr.Put{
			Token:    op,
			Children: childs,
		}
	}

	p.peekError(token.RIGHT_BRACE, "Missing statement end")
	p.advance()
	return stmt
}

func (p *Parser) parseArguments() expr.Node {
	var child expr.Node
	p.peekErrorMany("Missing or unknown argument",
		token.FLOAT,
		token.STRING,
		token.IDENT,
		token.BOOL,
		token.LEFT_CURLY,
		token.LEFT_BRACKET,
		token.TEMPLATE_STRING)
	if p.peekIs(token.LEFT_BRACKET) {
		child = p.parseIndex()
	} else if p.peekIs(token.LEFT_CURLY) {
		child = p.parseObject()
	} else if p.peekIs(token.TEMPLATE_STRING) {
		child = p.parseTemplateString()
	} else if p.peekIs(token.FLOAT) {
		t := p.peek()
		value, err := strconv.ParseFloat(t.Raw, 64)
		if err != nil {
			serror.Add(t, "Failed to parse number", "%q not a valid floating point integer", t.Raw)
			value = 0
		}
		child = &expr.Float{
			Token: t,
			Value: value,
		}
	} else if p.peekIs(token.STRING) {
		child = &expr.String{
			Token: p.peek(),
		}
	} else if p.peekIs(token.IDENT) {
		child = &expr.Ident{
			Token: p.peek(),
			Name:  p.peek().Raw,
		}
	} else if p.peekIs(token.BOOL) {
		child = &expr.Boolean{
			Token: p.peek(),
			// fastpath for easy boolean access, skipping a compare for each eval
			Value: p.peek().Raw == "true",
		}
	}
	return child
}

func (p *Parser) parseIndex() expr.Node {
	p.peekError(token.LEFT_BRACKET, "missing index start")
	o := expr.Index{
		Token: p.peek(),
		Index: make([]expr.Node, 0),
	}
	p.advance()
	p.peekError(token.IDENT, "missing element to index into")
	o.Target = p.parseArguments()
	p.advance()
	p.peekError(token.DOT, "missing index element and property divider")
	p.advance()
	o.Index = append(o.Index, p.parseArguments())
	p.advance()

	for p.peekIs(token.DOT) {
		p.advance()
		o.Index = append(o.Index, p.parseArguments())
		p.advance()
	}

	p.peekError(token.RIGHT_BRACKET, "missing index end")
	return &o
}

func (p *Parser) parseObject() expr.Node {
	p.peekError(token.LEFT_CURLY, "missing object start")
	o := expr.Object{
		Token:    p.peek(),
		Children: make([]expr.ObjectPair, 0),
	}
	p.advance()
	for !p.peekIs(token.RIGHT_CURLY) && !p.peekIs(token.EOF) {
		op := expr.ObjectPair{
			Key: p.parseArguments(),
		}
		p.advance()
		if p.peek().Type != token.COLON {
			p.peekError(token.COLON, "missing object key value divider")
			return nil
		}
		p.advance()
		op.Value = p.parseArguments()
		p.advance()
		o.Children = append(o.Children, op)
	}
	p.peekError(token.RIGHT_CURLY, "missing object end")
	return &o
}

func (p *Parser) parseTemplateString() *expr.TemplateString {
	t := &expr.TemplateString{
		Token:    p.peek(),
		Children: make([]expr.Node, 0),
	}
	p.advance()
	for !p.peekIs(token.TEMPLATE_STRING) {
		t.Children = append(t.Children, p.parseArguments())
		p.advance()
	}
	return t
}

func (p *Parser) advance() {
	if p.peek().Type == token.EOF {
		return
	}
	p.pos++
}

func (p *Parser) peek() *token.Token {
	return p.token[p.pos]
}

func (p *Parser) peekNext() *token.Token {
	if p.peekIs(token.EOF) {
		return p.peek()
	}
	return p.token[p.pos]
}

func (p *Parser) peekIs(tokenType int) bool {
	return p.peek().Type == tokenType
}

func (p *Parser) peekErrorMany(error string, tokenType ...int) {
	contains := false
	for _, t := range tokenType {
		if p.peekIs(t) {
			contains = true
		}
	}
	if !contains {
		o := make([]string, len(tokenType))
		for i, w := range tokenType {
			o[i] = token.TOKEN_NAME_MAP[w]
		}
		wanted := strings.Join(o, ",")
		t := p.peek()
		serror.Add(t, "Unexpected Token", "%s: Expected any of '%s' got '%s'.", error, wanted, token.TOKEN_NAME_MAP[t.Type])
	}
}

func (p *Parser) peekError(tokenType int, error string) (r bool) {
	if !p.peekIs(tokenType) {
		t := p.peek()
		serror.Add(t, "Unexpected Token", "%s: Expected Token '%s' got '%s'.", error, token.TOKEN_NAME_MAP[tokenType], token.TOKEN_NAME_MAP[t.Type])
		return true
	}
	return false
}
