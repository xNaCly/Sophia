package run

import (
	"fmt"
	"github.com/chzyer/readline"
	"log"
	"sophia/core"
	"sophia/core/consts"
	"strings"
)

func repl(run func(input []byte, filename string) ([]string, error)) {
	log.SetFlags(0)
	fmt.Println(`Welcome to the Sophia programming language repl - press <CTRL-D> or <CTRL-C> to quit...`)

	rl, err := readline.NewEx(&readline.Config{
		Prompt: "ß :: ",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

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
			val, error := run([]byte(line), "repl")
			if error != nil {
				log.Println(error)
			} else {
				fmt.Println("=", val)
			}
		}
	}
}
