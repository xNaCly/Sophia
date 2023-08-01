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

func run(input []byte) (s []string, e error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("err: %s", err)
			e = errors.New("runtime error")
			return
		}
	}()
	core.DbgLog("starting lexer")
	l := core.NewLexer(input)
	tokens := l.Lex()
	core.DbgLog("lexed", len(tokens), "token")
	if core.CONF.Debug {
		v, _ := json.MarshalIndent(tokens, "", "  ")
		core.DbgLog(string(v))
	}

	core.DbgLog("starting parser")
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

	if core.CONF.Debug {
		v, _ := json.MarshalIndent(ast, "", "  ")
		core.DbgLog(string(v))
	}
	core.DbgLog("done parsing - no errors, starting eval")
	s = core.Eval(ast)
	return
}

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	execute := flag.String("exp", "", "specifiy expression to execute")
	dbg := flag.Bool("dbg", false, "enable debug logs")
	flag.Parse()
	core.CONF.Debug = *dbg

	if len(*execute) != 0 {
		core.DbgLog("got -exp flag, running...")
		_, err := run([]byte(*execute))
		if err != nil {
			log.Fatalln(err)
		}
	} else if len(flag.Args()) == 1 {
		core.DbgLog("got file, running...")
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
		fmt.Print(core.ASCII_ART, "\n")
		core.DbgLog("go nothing, starting repl...")
		core.Repl(run)
	}
}
