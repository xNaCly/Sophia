package lexer

import (
	"sophia/core/serror"
	"sophia/core/token"
	"strings"
	"unicode"
)

type Lexer struct {
	input          []byte
	pos            int
	chr            byte
	line           int
	linepos        int
	errorFormatter *serror.ErrorFormatter
}

func New(input []byte, errorFormatter *serror.ErrorFormatter) *Lexer {
	if len(input) == 0 || len(strings.TrimSpace(string(input))) == 0 {
		errorFormatter.Add(nil, "Unexpected end of file", "Source possibly empty")
		return &Lexer{}
	}
	return &Lexer{
		input:          input,
		pos:            0,
		chr:            input[0],
		line:           0,
		linepos:        0,
		errorFormatter: errorFormatter,
	}
}

func (l *Lexer) Lex() []token.Token {
	t := make([]token.Token, 0)
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
			ttype = token.SUB
		case '/':
			ttype = token.DIV
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
			ttype = token.COLON
		case '.':
			ttype = token.DOT
		case '_':
			ttype = token.PARAM
		case '\'':
			t = append(t, l.templateString()...)
			continue
		case ' ', '\t', '\r', '\n':
			if l.chr == '\n' {
				l.linepos = 0
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
			if unicode.IsLetter(rune(l.chr)) {
				t = append(t, l.ident())
				continue
			} else if unicode.IsDigit(rune(l.chr)) {
				if tok, err := l.float(); err == nil {
					t = append(t, tok)
				} else {
					l.errorFormatter.Add(&tok, "Invalid floating point number", "")
				}
				continue
			}
		}

		if ttype == token.UNKNOWN {
			l.errorFormatter.Add(&token.Token{
				Pos:  l.pos,
				Type: ttype,
				Line: l.line,
				Raw:  string(l.chr),
			}, "Unknown character %q", "")
		}

		t = append(t, token.Token{
			Pos:  l.pos,
			Type: ttype,
			Line: l.line,
			Raw:  string(l.chr),
		})

		l.advance()
	}
	t = append(t, token.Token{
		Type: token.EOF, Line: l.line,
	})
	return t
}

func (l *Lexer) templateString() []token.Token {
	el := make([]token.Token, 0)
	el = append(el, token.Token{
		Type: token.TEMPLATE_STRING,
		Pos:  l.pos,
		Line: l.line,
		Raw:  "",
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
				el = append(el, token.Token{
					Pos:  l.pos - (len(b.String()) + 2),
					Type: token.STRING,
					Raw:  b.String(),
					Line: l.line,
				})
				b.Reset()
			}
			el = append(el, l.ident())
			continue
		} else if l.chr == '\n' || l.chr == 0 {
			var errEl token.Token
			if len(el) > 1 {
				errEl = el[len(el)-1]
			}
			l.errorFormatter.Add(&errEl, "Unexpected new line or end of file in template string", "Consider closing the template string via ' or omitting the inserted new line")
			return []token.Token{}
		} else if l.chr == '\'' {
			if b.Len() != 0 {
				el = append(el, token.Token{
					Pos:  l.pos - (len(b.String()) + 2),
					Type: token.STRING,
					Raw:  b.String(),
					Line: l.line,
				})
				b.Reset()
			}
			break
		}
		b.WriteByte(l.chr)
		l.advance()
	}

	if l.chr == '\'' {
		l.advance()
	}

	el = append(el, token.Token{
		Type: token.TEMPLATE_STRING,
		Pos:  l.pos,
		Line: l.line,
		Raw:  "",
	})
	return el
}

func (l *Lexer) string() token.Token {
	l.advance()
	b := strings.Builder{}
	for l.chr != '"' && l.chr != '\n' && l.chr != 0 {
		b.WriteByte(l.chr)
		l.advance()
	}
	str := b.String()
	if l.chr != '"' {
		l.errorFormatter.Add(&token.Token{
			Pos:  l.pos - (len(str) + 2),
			Type: token.STRING,
			Raw:  str,
			Line: l.line,
		}, "Unterminated string", "Consider closing the string via \"")
	} else {
		l.advance()
	}

	return token.Token{
		Pos:  l.pos - (len(str) + 2),
		Type: token.STRING,
		Raw:  str,
		Line: l.line,
	}
}

func (l *Lexer) ident() token.Token {
	builder := strings.Builder{}
	for unicode.IsLetter(rune(l.chr)) || l.chr == '_' || unicode.IsDigit(rune(l.chr)) {
		builder.WriteByte(l.chr)
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
	return token.Token{
		Pos:  l.pos - len(str),
		Type: ttype,
		Raw:  str,
		Line: l.line,
	}
}

func (l *Lexer) float() (token.Token, error) {
	builder := strings.Builder{}
	for unicode.IsDigit(rune(l.chr)) || l.chr == '.' || l.chr == '_' || l.chr == 'e' || l.chr == '-' {
		builder.WriteByte(l.chr)
		l.advance()
	}
	str := builder.String()
	return token.Token{
		Pos:  l.pos - len(str),
		Type: token.FLOAT,
		Raw:  str,
		Line: l.line,
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
