package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"sophia/core"
)

func run(input []byte) (f []float64, e error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("err: %s", err)
			e = errors.New("runtime error")
			return
		}
	}()
	l := core.NewLexer(input)
	tokens := l.Lex()

	p := core.NewParser(tokens)
	ast := p.Parse()
	if l.HasError {
		e = errors.New("lexer error")
		return
	}
	if p.HasError {
		e = errors.New("parser error")
		return
	}

	f = core.Eval(ast)
	return
}

func main() {
	log.SetFlags(log.Ltime)
	execute := flag.String("exp", "", "specifiy expression to execute")
	flag.Parse()

	if len(*execute) != 0 {
		_, err := run([]byte(*execute))
		if err != nil {
			log.Fatalln(err)
		}
	} else if len(flag.Args()) == 1 {
		file := flag.Args()[0]
		f, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to open file: %s\n", err)
		}
		_, err = run(f)
		if err != nil {
			log.Println(err)
			log.Fatalf("error in source file '%s' detected, stopping...", file)
		}
	} else {
		core.Repl(run)
	}
}
