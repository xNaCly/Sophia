package core

import (
	"log"
	"strings"
)

type Parser struct {
	token    []Token
	pos      int
	HasError bool
}

func NewParser(token []Token) Parser {
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

func (p *Parser) Parse() []Node {
	res := make([]Node, 0)
	for !p.peekIs(EOF) {
		res = append(res, p.parseStatment())
	}
	if p.HasError {
		return []Node{}
	}
	return res
}

func (p *Parser) parseStatment() Node {
	childs := make([]Node, 0)
	var stmt Node
	p.peekError(LEFT_BRACE, "Missing statement start")
	p.advance()
	p.peekErrorMany("Missing or unknown operator", EXPECTED_KEYWORDS...)
	op := p.peek()
	p.advance()

	for {
		var child Node
		if p.peekIs(RIGHT_BRACE) || p.peekIs(EOF) {
			// BUG: this does not always work
			break
		} else if p.peekIs(LEFT_BRACE) {
			child = p.parseStatment()
		} else {
			p.peekErrorMany("Missing or unknown argument", FLOAT, STRING)
			if p.peekIs(FLOAT) {
				child = &Float{
					Token: p.peek(),
				}
			} else if p.peekIs(STRING) {
				child = &String{
					Token: p.peek(),
				}
			}
		}

		childs = append(childs, child)
		p.advance()
	}

	switch op.Type {
	case ADD:
		stmt = &Add{
			Token:    op,
			Children: childs,
		}
	case SUB:
		stmt = &Sub{
			Token:    op,
			Children: childs,
		}
	case DIV:
		stmt = &Div{
			Token:    op,
			Children: childs,
		}
	case MUL:
		stmt = &Mul{
			Token:    op,
			Children: childs,
		}
	case PUT:
		stmt = &Put{
			Token:    op,
			Children: childs,
		}
	}

	p.peekError(RIGHT_BRACE, "Missing statement end")
	p.advance()
	return stmt
}

func (p *Parser) advance() {
	if p.peek().Type == EOF {
		return
	}
	p.pos++
}

func (p *Parser) peek() Token {
	return p.token[p.pos]
}

func (p *Parser) peekNext() Token {
	if p.peekIs(EOF) {
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
			o[i] = TOKEN_NAME_MAP[w]
		}
		wanted := strings.Join(o, ",")
		log.Printf("err: %s - Expected any of: '%s', got '%s' [l: %d:%d]", error, wanted, TOKEN_NAME_MAP[p.peek().Type], p.peek().Line, p.peek().Pos)
		p.HasError = true
	}
}

func (p *Parser) peekError(tokenType int, error string) {
	if !p.peekIs(tokenType) {
		log.Printf("err: %s - Expected Token '%s' got '%s' [l: %d:%d]", error, TOKEN_NAME_MAP[tokenType], TOKEN_NAME_MAP[p.peek().Type], p.peek().Line, p.peek().Pos)
		p.HasError = true
	}
}
