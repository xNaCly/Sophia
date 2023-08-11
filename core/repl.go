package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sophia/core/consts"
)

func Repl(run func(input []byte) ([]string, error)) {
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
			case "symbols":
				fmt.Printf("%#v\n", consts.SYMBOL_TABLE)
			case "funcs":
				fmt.Printf("%#v\n", consts.FUNC_TABLE)
			case "debug":
				CONF.Debug = !CONF.Debug
				log.Printf("toggled debug logging to='%t'", CONF.Debug)
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
