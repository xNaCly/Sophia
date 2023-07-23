package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sophia/core"
)

func run(input []byte) ([]float64, error) {
	l := core.NewLexer(input)
	tokens := l.Lex()
	p := core.NewParser(tokens)
	ast := p.Parse()
	if l.HasError {
		return []float64{}, errors.New("lexer error")
	}
	if p.HasError {
		return []float64{}, errors.New("parser error")
	}
	// TODO: temp, until eval is done
	res, _ := json.MarshalIndent(ast, "", "\t")
	fmt.Println(string(res))
	out := core.Eval(ast)
	return out, nil
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
		_, err = run(f)
		if err != nil {
			log.Fatalf("error in source file '%s' detected, stopping...", *file)
		}
	} else {
		core.Repl(run)
	}
}
