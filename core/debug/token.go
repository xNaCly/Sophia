package debug

import (
	"sophia/core/token"
	"strings"
)

func Token(tokenArr []token.Token) string {
	b := strings.Builder{}
	b.WriteRune('\n')
	for _, t := range tokenArr {
		b.WriteString(token.TOKEN_NAME_MAP[t.Type])
		b.WriteRune('\n')
	}
	return b.String()
}
