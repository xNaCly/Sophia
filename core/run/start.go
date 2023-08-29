package run

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sophia/core"
	"sophia/core/debug"
)

func Start() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	execute := flag.String("exp", "", "specifiy expression to execute")
	target := flag.String("target", "", "specifiy target to compile sophia to")
	dbg := flag.Bool("dbg", false, "enable debug logs")
	flag.Parse()
	core.CONF = core.Config{
		Debug:  *dbg,
		Target: *target,
	}

	if len(*execute) != 0 {
		debug.Log("got -exp flag, running...")
		_, err := run([]byte(*execute))
		if err != nil {
			log.Fatalln(err)
		}
	} else if len(flag.Args()) == 1 {
		debug.Log("got file, running...")
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
		if len(core.CONF.Target) > 0 {
			log.Fatalf("got compile target %q, but no file or expression to compile, exiting...", core.CONF.Target)
		}
		fmt.Print(core.ASCII_ART, "\n")
		debug.Log("go nothing, starting repl...")
		repl(run)
	}
}
