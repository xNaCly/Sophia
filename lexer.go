package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Lexer struct {
	input   []byte
	pos     int
	chr     byte
	line    int
	linepos int
}

func NewLexer(input []byte) Lexer {
	return Lexer{
		input:   input,
		pos:     0,
		chr:     input[0],
		line:    0,
		linepos: 0,
	}
}

func (l *Lexer) Lex() []Token {
	token := make([]Token, 0)
	for l.chr != 0 {
		ttype := UNKNOWN

		switch l.chr {
		case '[':
			ttype = LEFT_BRACE
		case ']':
			ttype = RIGHT_BRACE
		case ' ', '\t', '\r', '\n':
			if l.chr == '\n' {
				l.linepos = 0
				l.line++
			}
			l.advance()
			continue
		case '"':
			token = append(token, l.string())
			continue
		case ';':
			if l.peek() == ';' {
				for l.chr != '\n' && l.chr != 0 {
					l.advance()
				}
				continue
			}
		default:
			if unicode.IsLetter(rune(l.chr)) || l.chr == '_' {
				if t, err := l.ident(); err == nil {
					token = append(token, t)
					continue
				} else {
					l.error(1, t.Raw)
				}
			}
		}

		if ttype == UNKNOWN {
			l.error(0, "")
		}

		token = append(token, Token{
			Pos:  l.pos,
			Type: ttype,
		})

		l.advance()
	}
	return token
}

func (l *Lexer) error(errType uint, ident string) {
	switch errType {
	case 2:
		log.Printf("err: Unterminated String at [%d:%d-%d]", l.line, l.linepos-len(ident), l.linepos)
	case 1:
		log.Printf("err: UNKNOWN identifier '%s' at [%d:%d-%d]", ident, l.line, l.linepos-len(ident), l.linepos)
	default:
		log.Printf("err: UNKNOWN token '%c' at [%d:%d]", l.chr, l.line, l.linepos)
	}
	lines := strings.Split(string(l.input), "\n")
	if len(lines) > 1 {
		fmt.Printf("\n%.3d |\t%s\n%.3d |\t%s\n\t%s%c\n\n", l.line, lines[l.line-1], l.line+1, lines[l.line], strings.Repeat(" ", l.linepos), '^')
	} else {
		fmt.Printf("\n%.3d |\t%s\n\t%s%c\n\n", l.line+1, lines[l.line], strings.Repeat(" ", l.linepos), '^')
	}
	os.Exit(1)
}

func (l *Lexer) string() Token {
	l.advance()
	b := strings.Builder{}
	for l.chr != '"' && l.chr != '\n' && l.chr != 0 {
		b.WriteByte(l.chr)
		l.advance()
	}
	str := b.String()
	if l.chr != '"' {
		l.error(2, str)
	} else {
		l.advance()
	}

	return Token{
		Pos:  l.pos - len(str),
		Type: STRING,
		Raw:  str,
	}
}

func (l *Lexer) ident() (Token, error) {
	builder := strings.Builder{}
	for unicode.IsLetter(rune(l.chr)) || l.chr == '_' {
		builder.WriteByte(l.chr)
		l.advance()
	}
	str := builder.String()
	if t, ok := KEYWORDS[str]; ok {
		return Token{
			Pos:  l.pos - len(str),
			Type: t,
			Raw:  str,
		}, nil
	}
	return Token{
		Raw: str,
	}, errors.New("failed to find identifier in keywords")
}

// TODO:
func (l *Lexer) float() Token {
	return Token{}
}

func (l *Lexer) peek() byte {
	if l.pos+1 < len(l.input) {
		return l.input[l.pos+1]
	}
	return 0
}

func (l *Lexer) advance() {
	if l.pos+1 < len(l.input) {
		l.pos++
		l.linepos++
		l.chr = l.input[l.pos]
		return
	}
	l.chr = 0
}
