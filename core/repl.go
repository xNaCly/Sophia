package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Repl(run func(input []byte) ([]float64, error)) {
	fmt.Println(`Welcome to the Tisp repl - press <CTRL-D> or <CTRL-C> to quit...`)
	prompt := "τ :: "
	scanner := bufio.NewScanner(os.Stdin)
	var last []float64
	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Bytes()
		switch string(line) {
		case ":last":
			fmt.Println("=", last)
			continue
		case ":quit":
			break
		}
		val, error := run(line)
		if error != nil {
			log.Println("err: error in input")
		} else {
			fmt.Println("=", val)
		}
		last = val
	}
}
