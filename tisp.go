package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"tisp/core"
)

func main() {
	log.SetFlags(log.Ltime)
	file := flag.String("f", "", "specifiy source file, if not specifiy start repl")
	flag.Parse()

	if len(*file) != 0 {
		f, err := os.ReadFile(*file)
		if err != nil {
			log.Fatalf("Failed to open file: %s\n", err)
		}

		l := core.NewLexer(f)
		tokens := l.Lex()

		if l.HasError {
			log.Fatalf("error in source file '%s' detected, stopping...", *file)
		}

		v, _ := json.MarshalIndent(tokens, "", "\t")
		log.Println(string(v))
	} else {
		core.Repl()
	}
}
