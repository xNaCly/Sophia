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
	stmt := &Statement{
		Children: make([]Node, 0),
	}
	p.peekError(LEFT_BRACE, "Missing statement start")
	p.advance()
	p.peekErrorMany(EXPECTED_KEYWORDS...)
	stmt.Token = p.peek()
	p.advance()

	for {
		var child Node
		if p.peekIs(RIGHT_BRACE) {
			break
		} else if p.peekIs(EOF) {
			return stmt
		}

		if p.peekIs(LEFT_BRACE) {
			child = p.parseStatment()
		} else {
			p.peekErrorMany(FLOAT, STRING)
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
		stmt.Children = append(stmt.Children, child)
		p.advance()
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

func (p *Parser) peekErrorMany(tokenType ...int) {
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
		log.Printf("err: Expected any of: '%s', got '%s' [l: %d:%d]", wanted, TOKEN_NAME_MAP[p.peek().Type], p.peek().Line, p.peek().Pos)
		p.HasError = true
	}
}

func (p *Parser) peekError(tokenType int, error string) {
	if !p.peekIs(tokenType) {
		log.Printf("err: Expected Token '%s' got '%s' - %s [l: %d:%d]", TOKEN_NAME_MAP[tokenType], TOKEN_NAME_MAP[p.peek().Type], error, p.peek().Line, p.peek().Pos)
		p.HasError = true
	}
}
