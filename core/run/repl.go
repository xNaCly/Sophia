package run

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"
	"github.com/xnacly/sophia/core"
	"github.com/xnacly/sophia/core/consts"
	"github.com/xnacly/sophia/core/serror"
)

func repl(run func(r io.Reader, filename string) ([]string, error)) {
	log.SetFlags(0)
	fmt.Println(`Welcome to the Sophia programming language repl - press <CTRL-D> or <CTRL-C> to quit...`)

	rl, err := readline.NewEx(&readline.Config{
		Prompt: "ß > ",
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
			var r io.Reader
			r = strings.NewReader(line)
			serror.SetDefault(serror.NewFormatter(&core.CONF, line, "repl", nil))
			val, error := run(r, "repl")
			if error != nil {
				log.Println(error)
			} else {
				fmt.Println("=")
				for _, v := range val {
					fmt.Println(" ", v)
				}
			}
		}
	}
}
