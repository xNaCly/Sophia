package debug

import (
	"fmt"
	"sophia/core/token"
	"strings"
)

// populates a string with the useful information contained in the tokenArray
func Token(tokenArr []token.Token) string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("\n%15s | %25s | %5s | %5s\n", "type", "raw", "line", "pos"))
	for _, t := range tokenArr {
		b.WriteString(fmt.Sprintf("%15s | %25s | %5d | %5d\n", token.TOKEN_NAME_MAP[t.Type], t.Raw, t.Line, t.Pos))
	}
	return b.String()
}
