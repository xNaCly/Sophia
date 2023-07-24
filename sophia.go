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

var DEBUG = false

func run(input []byte) ([]float64, error) {
	l := core.NewLexer(input)
	tokens := l.Lex()

	if DEBUG {
		tl := len(tokens)
		for i, t := range tokens {
			fmt.Printf("dbg: [%d/%d] %s at l=%d:p=%d with '%v'\n", i+1, tl, core.TOKEN_NAME_MAP[t.Type], t.Line, t.Pos, t.Raw)
		}
	}

	p := core.NewParser(tokens)
	ast := p.Parse()
	if DEBUG {
		v, _ := json.MarshalIndent(ast, "", "\t")
		fmt.Printf("dbg: %s", v)
	}
	if l.HasError {
		return []float64{}, errors.New("lexer error")
	}
	if p.HasError {
		return []float64{}, errors.New("parser error")
	}

	out := core.Eval(ast)
	return out, nil
}

func main() {
	log.SetFlags(log.Ltime)
	execute := flag.String("exp", "", "specifiy expression to execute")
	flag.BoolVar(&DEBUG, "dbg", false, "enable debug mode, prints lexing, parsing and eval information as well as timestamps")
	flag.Parse()

	if len(*execute) != 0 {
		run([]byte(*execute))
	} else if len(os.Args) > 1 {
		file := os.Args[1]
		f, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to open file: %s\n", err)
		}
		_, err = run(f)
		if err != nil {
			log.Fatalf("error in source file '%s' detected, stopping...", file)
		}
	} else {
		core.Repl(run)
	}
}
