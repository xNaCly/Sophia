package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Repl() {
	fmt.Println(`Welcome to the Tisp repl - press <CTRL-D> or <CTRL-C> to quit...`)
	prompt := "τ :: "
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Bytes()
		l := NewLexer(line)
		v, _ := json.MarshalIndent(l.Lex(), "", "\t")
		log.Println(string(v))
	}
}
