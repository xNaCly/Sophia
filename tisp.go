package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"
	"tisp/core"
)

func run(input []byte) error {
	l := core.NewLexer(input)
	tokens := l.Lex()
	p := core.NewParser(tokens)
	ast := p.Parse()
	if l.HasError {
		return errors.New("lexer error")
	}
	if p.HasError {
		return errors.New("parser error")
	}

	v, _ := json.MarshalIndent(ast, "", "\t")
	log.Printf("%s\n", v)
	return nil
}

func main() {
	log.SetFlags(log.Ltime)
	file := flag.String("f", "", "specifiy source file, if not specifiy start repl")
	execute := flag.String("e", "", "specifiy expression to execute")
	flag.Parse()

	if len(*execute) != 0 {
		run([]byte(*execute))
	} else if len(*file) != 0 {
		f, err := os.ReadFile(*file)
		if err != nil {
			log.Fatalf("Failed to open file: %s\n", err)
		}
		err = run(f)
		if err != nil {
			log.Fatalf("error in source file '%s' detected, stopping...", *file)
		}
	} else {
		core.Repl(run)
	}
}
