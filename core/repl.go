package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Repl(run func(input []byte) ([]float64, error)) {
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

		val, error := run(line)
		if error != nil {
			log.Println(error)
		} else {
			fmt.Println("=", val)
		}
	}
}
