package debug

import (
	"fmt"
	"os"
	"sophia/core"
)

const (
	ANSI_RESET = "\033[0m"
	ANSI_BLUE  = "\033[94m"
)

func Log(in ...any) {
	if core.CONF.Debug {
		os.Stdout.WriteString(ANSI_BLUE)
		os.Stdout.WriteString("info: ")
		os.Stdout.WriteString(ANSI_RESET)
		fmt.Println(in...)
	}
}
