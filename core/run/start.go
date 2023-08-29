package run

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sophia/core"
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
		repl(run)
	}
}
