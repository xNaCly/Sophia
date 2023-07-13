package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"tisp/core"
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

		l := core.NewLexer(f)
		tokens := l.Lex()

		v, _ := json.MarshalIndent(tokens, "", "\t")
		log.Println(string(v))
	} else {
		core.Repl()
	}
}
