package debug

import (
	"fmt"
	"os"
	"sophia/core"
	"time"
)

const (
	ANSI_RESET = "\033[0m"
	ANSI_BLUE  = "\033[94m"
)

func Log(in ...any) {
	if core.CONF.Debug {
		fmt.Print(time.Now().Format("15:04:05.000000000"), " ")
		os.Stdout.WriteString(ANSI_BLUE)
		os.Stdout.WriteString("info: ")
		os.Stdout.WriteString(ANSI_RESET)
		fmt.Println(in...)
	}
}

func Logf(format string, in ...any) {
	if core.CONF.Debug {
		fmt.Print(time.Now().Format("15:04:05.000000000"), " ")
		os.Stdout.WriteString(ANSI_BLUE)
		os.Stdout.WriteString("info: ")
		os.Stdout.WriteString(ANSI_RESET)
		fmt.Printf(format, in...)
	}
}
