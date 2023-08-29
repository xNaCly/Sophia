package run

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sophia/core"
	"sophia/core/consts"
)

func repl(run func(input []byte) ([]string, error)) {
	fmt.Println(`Welcome to the Sophia repl - press <CTRL-D> or <CTRL-C> to quit...`)
	prompt := "ß :: "
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Bytes()
		if line[0] == '~' {
			switch string(line[1:]) {
			case "syms":
				fmt.Printf("%#v\n", consts.SYMBOL_TABLE)
			case "funs":
				fmt.Printf("%#v\n", consts.FUNC_TABLE)
			case "debug":
				core.CONF.Debug = !core.CONF.Debug
				log.Printf("toggled debug logging to='%t'", core.CONF.Debug)
			}
		} else {
			val, error := run(line)
			if error != nil {
				log.Println(error)
			} else {
				fmt.Println("=", val)
			}
		}
	}
}
