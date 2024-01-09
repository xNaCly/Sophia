package builtin

import (
	"os"
	"github.com/xnacly/sophia/core/shared"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
	"strings"
)

var sharedPrintBuffer = &strings.Builder{}

func builtinPrintln(tok *token.Token, args ...types.Node) any {
	sharedPrintBuffer.Reset()
	shared.FormatHelper(sharedPrintBuffer, args, ' ')
	sharedPrintBuffer.WriteRune('\n')
	os.Stdout.WriteString(sharedPrintBuffer.String())
	return nil
}
