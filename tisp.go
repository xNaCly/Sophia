package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)
	file := flag.String("f", "", "specifiy if following argument should be considered")
	flag.Parse()

	if len(*file) != 0 {
		f, err := os.ReadFile(*file)
		if err != nil {
			log.Fatalf("Failed to open file: %s\n", err)
		}

		l := NewLexer(f)
		tokens := l.Lex()

		v, _ := json.MarshalIndent(tokens, "", "\t")
		log.Println(string(v))
	} else {
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
}
