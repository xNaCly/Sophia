package core

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Lexer struct {
	input    []byte
	pos      int
	chr      byte
	line     int
	linepos  int
	HasError bool
}

func NewLexer(input []byte) Lexer {
	if len(input) == 0 || len(strings.TrimSpace(string(input))) == 0 {
		log.Println("err: input is empty, stopping")
		return Lexer{
			HasError: true,
		}
	}
	return Lexer{
		input:    input,
		pos:      0,
		chr:      input[0],
		line:     0,
		linepos:  0,
		HasError: false,
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
				} else {
					l.error(1, t.Raw)
				}
				continue
			} else if unicode.IsDigit(rune(l.chr)) || l.chr == '.' {
				if t, err := l.float(); err == nil {
					token = append(token, t)
				} else {
					l.error(3, t.Raw)
				}
				continue
			}
		}

		if ttype == UNKNOWN {
			l.error(0, "")
		}

		token = append(token, Token{
			Pos:  l.pos,
			Type: ttype,
			Line: l.line,
		})

		l.advance()
	}
	if l.HasError {
		return []Token{
			{Type: EOF, Line: l.line},
		}
	}
	token = append(token, Token{
		Type: EOF, Line: l.line,
	})
	return token
}

func (l *Lexer) error(errType uint, ident string) {
	pos := l.linepos
	iLen := len(ident)

	switch errType {
	case 3:
		pos = pos - iLen
		log.Printf("err: Invalid floating point integer '%s' at [l %d:%d->%d]", ident, l.line+1, pos, l.linepos-1)
	case 2:
		pos = pos - iLen
		log.Printf("err: Unterminated String at [l %d:%d->%d]", l.line+1, pos, l.linepos-1)
	case 1:
		pos = pos - iLen
		log.Printf("err: UNKNOWN identifier '%s' at [l %d:%d->%d]", ident, l.line+1, pos, l.linepos-1)
		pos++
	default:
		log.Printf("err: UNKNOWN token '%c' at [l %d:%d]", l.chr, l.line+1, pos)
	}

	lines := strings.Split(string(l.input), "\n")

	spaces := pos

	if iLen > 0 {
		spaces = pos - iLen
	}

	if spaces < 0 {
		spaces = 0
	}

	// if string error highlight string start " and predicted end " with ^
	if errType == 2 {
		iLen++
	}

	// if no identifier given, print one ^
	if iLen == 0 {
		iLen += 1
	}

	if l.line-1 > -1 {
		spaces -= 1
		if spaces < 0 {
			spaces = 0
		}
		fmt.Printf("\n%.3d |\t%s\n%.3d |\t%s\n\t%s%s\n\n", l.line, lines[l.line-1], l.line+1, lines[l.line], strings.Repeat(" ", spaces), strings.Repeat("^", iLen))
	} else {
		fmt.Printf("\n%.3d |\t%s\n\t%s%s\n\n", l.line+1, lines[l.line], strings.Repeat(" ", spaces), strings.Repeat("^", iLen))
	}
	l.HasError = true
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
		Line: l.line,
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
			Line: l.line,
		}, nil
	}
	return Token{
		Raw:  str,
		Line: l.line,
	}, errors.New("failed to find identifier in keywords")
}

func (l *Lexer) float() (Token, error) {
	builder := strings.Builder{}
	for unicode.IsDigit(rune(l.chr)) || l.chr == '.' || l.chr == '_' || l.chr == 'e' || l.chr == '-' {
		builder.WriteByte(l.chr)
		l.advance()
	}
	str := builder.String()
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return Token{
			Raw:  str,
			Line: l.line,
		}, err
	}
	return Token{
		Pos:   l.pos - len(str),
		Type:  FLOAT,
		Raw:   str,
		Float: float,
		Line:  l.line,
	}, nil
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
