package lexer

import (
	"bufio"
	"io"
	"strings"
	"unicode"

	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
)

type Lexer struct {
	reader  *bufio.Reader
	pos     int
	chr     rune
	line    int
	linepos int
}

func New(r io.Reader) *Lexer {
	in := bufio.NewReader(r)
	if in.Size() == 0 {
		serror.Add(&token.Token{LinePos: 0, Raw: " "}, "Unexpected end of file", "Source empty")
		return &Lexer{}
	}

	l := &Lexer{
		reader:  in,
		pos:     0,
		line:    0,
		linepos: 0,
	}
	l.advance()
	return l
}

func (l *Lexer) Lex() []*token.Token {
	t := make([]*token.Token, 0)
	for l.chr != 0 {
		ttype := token.UNKNOWN

		switch l.chr {
		case '+':
			if l.peek() == '+' {
				ttype = token.MERGE
				l.advance()
			} else {
				ttype = token.ADD
			}
		case '-':
			if unicode.IsDigit(l.peek()) {
				if tok, err := l.float(); err == nil {
					t = append(t, tok)
				} else {
					serror.Add(tok, "Invalid floating point number", "")
				}
				continue
			} else {
				ttype = token.SUB
			}
		case '/':
			ttype = token.DIV
		case '#':
			ttype = token.HASHTAG
		case '*':
			ttype = token.MUL
		case '%':
			ttype = token.MOD
		case '(':
			ttype = token.LEFT_BRACE
		case ')':
			ttype = token.RIGHT_BRACE
		case '{':
			ttype = token.LEFT_CURLY
		case '}':
			ttype = token.RIGHT_CURLY
		case '[':
			ttype = token.LEFT_BRACKET
		case ']':
			ttype = token.RIGHT_BRACKET
		case ':':
			if l.peek() == ':' {
				ttype = token.DOUBLE_COLON
				l.advance()
			} else {
				ttype = token.COLON
			}
		case '.':
			ttype = token.DOT
		case '=':
			ttype = token.EQUAL
		case '<':
			ttype = token.LT
		case '>':
			ttype = token.GT
		case '\'':
			t = append(t, l.templateString()...)
			continue
		case ' ', '\t', '\r', '\n':
			if l.chr == '\n' {
				l.linepos = -1
				l.line++
			}
			l.advance()
			continue
		case '"':
			t = append(t, l.string())
			continue
		case ';':
			if l.peek() == ';' {
				for l.chr != '\n' && l.chr != 0 {
					l.advance()
				}
				continue
			}
		default:
			if unicode.IsLetter(l.chr) {
				t = append(t, l.ident())
				continue
			} else if unicode.IsDigit(l.chr) {
				if tok, err := l.float(); err == nil {
					t = append(t, tok)
				} else {
					serror.Add(tok, "Invalid floating point number", "")
				}
				continue
			}
		}

		if ttype == token.UNKNOWN {
			serror.Add(&token.Token{
				Pos:     l.pos,
				Type:    ttype,
				Line:    l.line,
				LinePos: l.linepos,
				Raw:     string(l.chr),
			}, "Unknown character", "Unexpected \"%c\"", l.chr)
		}

		t = append(t, &token.Token{
			Pos:     l.pos,
			Type:    ttype,
			Line:    l.line,
			LinePos: l.linepos,
			Raw:     string(l.chr),
		})

		l.advance()
	}
	t = append(t, &token.Token{
		Type:    token.EOF,
		Line:    l.line,
		LinePos: l.linepos,
		Raw:     " ",
	})
	return t
}

func (l *Lexer) templateString() []*token.Token {
	el := make([]*token.Token, 0)
	el = append(el, &token.Token{
		Type:    token.TEMPLATE_STRING,
		LinePos: l.linepos,
		Pos:     l.pos,
		Line:    l.line,
		Raw:     "",
	})
	b := strings.Builder{}

	l.advance() // skip '

	for {
		if l.chr == '}' {
			l.advance()
			continue
		} else if l.chr == '{' {
			l.advance()
			if b.Len() != 0 {
				el = append(el, &token.Token{
					Pos:     l.pos - (len(b.String()) + 2),
					Type:    token.STRING,
					Raw:     b.String(),
					Line:    l.line,
					LinePos: l.linepos,
				})
				b.Reset()
			}
			el = append(el, l.ident())
			continue
		} else if l.chr == '\n' || l.chr == 0 {
			var errEl *token.Token
			if len(el) > 1 {
				errEl = el[len(el)-1]
			}
			serror.Add(errEl, "Unexpected new line or end of file in template string", "Consider closing the template string via ' or omitting the inserted new line")
			return []*token.Token{}
		} else if l.chr == '\'' {
			if b.Len() != 0 {
				el = append(el, &token.Token{
					Pos:     l.pos - (len(b.String()) + 2),
					LinePos: l.linepos,
					Type:    token.STRING,
					Raw:     b.String(),
					Line:    l.line,
				})
				b.Reset()
			}
			break
		}
		b.WriteRune(l.chr)
		l.advance()
	}

	if l.chr == '\'' {
		l.advance()
	}

	el = append(el, &token.Token{
		Type:    token.TEMPLATE_STRING,
		LinePos: l.linepos,
		Pos:     l.pos,
		Line:    l.line,
		Raw:     "",
	})
	return el
}

func (l *Lexer) string() *token.Token {
	l.advance()
	b := strings.Builder{}
	for l.chr != '"' && l.chr != 0 {
		if l.chr == '\n' {
			l.line++
			l.linepos = -1
		}
		b.WriteRune(l.chr)
		l.advance()
	}
	str := b.String()
	if l.chr != '"' {
		serror.Add(&token.Token{
			Pos:     l.pos - (len(str) + 2),
			Type:    token.STRING,
			Raw:     "\"" + str,
			Line:    l.line,
			LinePos: l.linepos - len(str),
		}, "Unterminated string", "Consider closing the string via \"")
	} else {
		l.advance()
	}

	return &token.Token{
		Pos:     l.pos - (len(str) + 2),
		Type:    token.STRING,
		LinePos: l.linepos - len(str) - 1,
		Raw:     str,
		Line:    l.line,
	}
}

func (l *Lexer) ident() *token.Token {
	builder := strings.Builder{}
	for unicode.IsLetter(l.chr) || l.chr == '_' || unicode.IsDigit(l.chr) || l.chr == '-' {
		builder.WriteRune(l.chr)
		l.advance()
	}
	str := builder.String()
	ttype := token.UNKNOWN
	switch str {
	case "true", "false":
		ttype = token.BOOL
	default:
		ttype = token.IDENT
	}
	if tokenType, ok := token.KEYWORD_MAP[str]; ok {
		ttype = tokenType
	}
	return &token.Token{
		Pos:     l.pos - len(str),
		Type:    ttype,
		LinePos: l.linepos - len(str),
		Raw:     str,
		Line:    l.line,
	}
}

func (l *Lexer) float() (*token.Token, error) {
	builder := strings.Builder{}
	for unicode.IsDigit(l.chr) || (l.chr == '.' && unicode.IsDigit(l.peek())) || l.chr == '_' || l.chr == 'e' || l.chr == '-' {
		builder.WriteRune(l.chr)
		l.advance()
	}
	str := builder.String()
	return &token.Token{
		Pos:     l.pos - len(str),
		Type:    token.FLOAT,
		LinePos: l.linepos - len(str),
		Raw:     str,
		Line:    l.line,
	}, nil
}

func (l *Lexer) peek() rune {
	cc, _, err := l.reader.ReadRune()
	if err != nil {
		return 0
	}
	l.reader.UnreadRune()
	return cc
}

func (l *Lexer) advance() {
	cc, _, err := l.reader.ReadRune()
	if err != nil {
		l.chr = 0
		return
	}
	l.pos++
	l.linepos++
	l.chr = cc
}
