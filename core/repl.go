package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Repl(run func(input []byte) error) {
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
		error := run(line)
		if error != nil {
			log.Println("err: error in input")
		}
	}
}
