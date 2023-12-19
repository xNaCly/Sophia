package builtin

import (
	"os"
	"sophia/core/shared"
	"sophia/core/token"
	"sophia/core/types"
	"strings"
)

var sharedPrintBuffer = &strings.Builder{}

func buildinPrintln(tok *token.Token, args ...types.Node) any {
	sharedPrintBuffer.Reset()
	shared.FormatHelper(sharedPrintBuffer, args, ' ')
	sharedPrintBuffer.WriteRune('\n')
	os.Stdout.WriteString(sharedPrintBuffer.String())
	return nil
}
