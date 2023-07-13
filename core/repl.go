package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Repl() {
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
