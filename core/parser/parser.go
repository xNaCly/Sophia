package parser

import (
	"io"
	"log"
	"os"
	"sophia/core/expr"
	"sophia/core/lexer"
	"sophia/core/token"
	"strings"
)

// TODO: error display like in the lexer

type Parser struct {
	token    []token.Token
	filename string
	pos      int
	HasError bool
}

func New(token []token.Token, filename string) Parser {
	if len(token) == 0 {
		log.Println("Parser got no tokens, stopping...")
		return Parser{
			HasError: true,
		}
	}
	return Parser{
		token:    token,
		pos:      0,
		filename: filename,
		HasError: false,
	}
}

func (p *Parser) Parse() []expr.Node {
	res := make([]expr.Node, 0)
	for !p.peekIs(token.EOF) {
		stmt := p.parseStatment()
		if stmt.GetToken().Type == token.LOAD {
			if loadStmt, ok := stmt.(*expr.Load); ok {
				res = append(res, p.loadNewSource(loadStmt)...)
			}
			continue
		}
		res = append(res, stmt)
	}
	if p.HasError {
		return []expr.Node{}
	}
	return res
}

func (p *Parser) loadNewSource(node *expr.Load) []expr.Node {
	res := make([]expr.Node, 0)
	for i := 0; i < len(node.Imports); i++ {
		name := node.Imports[i]
		file, err := os.Open(name)
		if err != nil {
			log.Panicf("failed to open %q: %s", name, err)
			p.HasError = true
			return nil
		}
		content, err := io.ReadAll(file)
		if err != nil {
			log.Panicf("failed to read from %q: %s", name, err)
			p.HasError = true
			return nil
		}
		lexer := lexer.New(content)
		token := lexer.Lex()
		if name == p.filename {
			log.Panicf("detected recursion in file imports, got %q while already parsing %q", name, p.filename)
			p.HasError = true
			return nil
		}
		parser := New(token, name)
		res = append(res, parser.Parse()...)
	}
	return res
}

func (p *Parser) parseArguments() expr.Node {
	var child expr.Node
	p.peekErrorMany("Missing or unknown argument", token.FLOAT, token.STRING, token.IDENT, token.BOOL, token.TEMPLATE_STRING)
	if p.peekIs(token.TEMPLATE_STRING) {
		child = p.parseTemplateString()
	} else if p.peekIs(token.FLOAT) {
		child = &expr.Float{
			Token: p.peek(),
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
		}
	}
	return child
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
			childs = append(childs, p.parseStatment())
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
	case token.MATCH:
		stmt = &expr.Match{
			Token:    op,
			Branches: childs,
		}
	case token.LOAD:
		if len(childs) == 0 {
			log.Printf("err: expected at least one argument for loading sources, got %d", len(childs))
			p.HasError = true
			return nil
		}
		imports := make([]string, len(childs))
		for i, c := range childs {
			if c.GetToken().Type != token.STRING {
				log.Printf("err: expected strings as load arguments, got %q", token.TOKEN_NAME_MAP[c.GetToken().Type])
				p.HasError = true
				return nil
			}
			imports[i] = c.GetToken().Raw
		}
		stmt = &expr.Load{
			Token:   op,
			Imports: imports,
		}
	case token.FOR:
		if len(childs) < 2 {
			log.Printf("err: expected two argument for loop definition, got %d", len(childs))
			p.HasError = true
			return nil
		}
		params := childs[0]
		if params.GetToken().Type != token.PARAM {
			log.Printf("err: expected the first argument for loop definition to be of type PARAM, got %q", token.TOKEN_NAME_MAP[childs[0].GetToken().Type])
			p.HasError = true
			return nil
		}
		if len(params.(*expr.Params).Children) != 1 {
			log.Printf("err: expected one parameter for loop definition, got %d", len(params.(*expr.Params).Children))
			p.HasError = true
			return nil
		}
		// TODO: check if first is of type params
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
			log.Printf("err: expected exactly two statements for less than comparison, got %d", len(childs))
			p.HasError = true
			return nil
		}
		stmt = &expr.Lt{
			Token:    op,
			Children: childs,
		}
	case token.GT:
		if len(childs) != 2 {
			log.Printf("err: expected exactly two statements for greater than comparison, got %d", len(childs))
			p.HasError = true
			return nil
		}
		stmt = &expr.Gt{
			Token:    op,
			Children: childs,
		}
	case token.PARAM:
		for _, c := range childs {
			t := c.GetToken().Type
			if t != token.IDENT {
				log.Printf("err: expected identifier for parameter definition, got %q", token.TOKEN_NAME_MAP[t])
				p.HasError = true
				return nil
			}
		}
		stmt = &expr.Params{
			Token:    op,
			Children: childs,
		}
	case token.FUNC:
		if len(childs) < 2 {
			log.Printf("err: expected at least two argument for function definition, got %d", len(childs))
			p.HasError = true
			return nil
		}
		if childs[0].GetToken().Type != token.IDENT {
			log.Printf("err: expected the first argument for function definition to be of type IDENT, got %q", token.TOKEN_NAME_MAP[childs[0].GetToken().Type])
			p.HasError = true
			return nil
		}
		if childs[1].GetToken().Type != token.PARAM {
			log.Printf("err: expected the second argument for function definition to be of type PARAM, got %q", token.TOKEN_NAME_MAP[childs[0].GetToken().Type])
			p.HasError = true
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
			log.Printf("err: expected at least two argument for condition, got %d", len(childs))
			p.HasError = true
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
			log.Printf("err: expected at least one argument for variable declaration, got %d", len(childs))
			p.HasError = true
			return nil
		}
		ident := childs[0]
		if ident.GetToken().Type != token.IDENT {
			log.Printf("err: expected 'IDENT' as first argument in variable declaration, got %s", token.TOKEN_NAME_MAP[ident.GetToken().Type])
			p.HasError = true
			return nil
		}
		stmt = &expr.Var{
			Token: op,
			Name:  ident.GetToken().Raw,
			Value: childs[1:],
		}
	case token.MERGE:
		stmt = &expr.Merge{
			Token:    op,
			Children: childs,
		}
	case token.EQUAL:
		stmt = &expr.Equal{
			Token:    op,
			Children: childs,
		}
	case token.NEG:
		if len(childs) != 1 {
			log.Printf("err: expected exactly one argument for negation, got %d", len(childs))
			p.HasError = true
			return nil
		}
		stmt = &expr.Neg{
			Token:    op,
			Children: childs[0],
		}
	case token.OR:
		stmt = &expr.Or{
			Token:    op,
			Children: childs,
		}
	case token.AND:
		stmt = &expr.And{
			Token:    op,
			Children: childs,
		}
	case token.ADD:
		stmt = &expr.Add{
			Token:    op,
			Children: childs,
		}
	case token.SUB:
		stmt = &expr.Sub{
			Token:    op,
			Children: childs,
		}
	case token.DIV:
		stmt = &expr.Div{
			Token:    op,
			Children: childs,
		}
	case token.MUL:
		stmt = &expr.Mul{
			Token:    op,
			Children: childs,
		}
	case token.MOD:
		stmt = &expr.Mod{
			Token:    op,
			Children: childs,
		}
	case token.PUT:
		stmt = &expr.Put{
			Token:    op,
			Children: childs,
		}
	}

	p.peekError(token.RIGHT_BRACE, "Missing statement end")
	p.advance()
	return stmt
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

func (p *Parser) peek() token.Token {
	return p.token[p.pos]
}

func (p *Parser) peekNext() token.Token {
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
		log.Printf("err: %s - Expected any of: '%s', got '%s' [l: %d:%d]", error, wanted, token.TOKEN_NAME_MAP[p.peek().Type], p.peek().Line, p.peek().Pos)
		p.HasError = true
	}
}

func (p *Parser) peekError(tokenType int, error string) {
	if !p.peekIs(tokenType) {
		log.Printf("err: %s - Expected Token '%s' got '%s' [l: %d:%d]", error, token.TOKEN_NAME_MAP[tokenType], token.TOKEN_NAME_MAP[p.peek().Type], p.peek().Line, p.peek().Pos)
		p.HasError = true
	}
}
