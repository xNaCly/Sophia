package parser

import (
	"log"
	"sophia/core/expr"
	"sophia/core/token"
	"strings"
)

// TODO: error display like in the lexer

type Parser struct {
	token    []token.Token
	pos      int
	HasError bool
}

func New(token []token.Token) Parser {
	if len(token) == 0 {
		log.Println("Parser got no tokens, stopping...")
		return Parser{
			HasError: true,
		}
	}
	return Parser{
		token:    token,
		pos:      0,
		HasError: false,
	}
}

func (p *Parser) Parse() []expr.Node {
	res := make([]expr.Node, 0)
	for !p.peekIs(token.EOF) {
		res = append(res, p.parseStatment())
	}
	if p.HasError {
		return []expr.Node{}
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
			childs = append(childs, p.parseStatment())
			continue
		} else {
			p.peekErrorMany("Missing or unknown argument", token.FLOAT, token.STRING, token.IDENT, token.BOOL)
			if p.peekIs(token.FLOAT) {
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
		}

		if child != nil {
			childs = append(childs, child)
		}
		p.advance()
	}

	switch op.Type {
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
	case token.COLON:
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
	case token.CONCAT:
		stmt = &expr.Concat{
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
