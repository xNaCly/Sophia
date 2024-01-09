package run

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xnacly/sophia/core"
	"github.com/xnacly/sophia/core/debug"
)

func Start() {
	execute := flag.String("exp", "", "specifiy expression to execute")
	dbg := flag.Bool("dbg", false, "enable debug logs")
	allErrors := flag.Bool("all-errors", false, "display all found errors")
	flag.Parse()
	core.CONF = core.Config{
		Debug:     *dbg,
		AllErrors: *allErrors,
	}

	if *dbg {
		log.SetFlags(log.Ltime | log.Lmicroseconds)
	} else {
		log.SetFlags(0)
	}

	stdinInf, err := os.Stdin.Stat()
	// check if stdin is readable and the process is in a pipe
	if err == nil && !(stdinInf.Mode()&os.ModeNamedPipe == 0) {
		debug.Log("got stdin content, running...")
		_, err = run(os.Stdin, "stdin")
		if err != nil {
			log.Fatalln(err)
		}
	} else if len(*execute) != 0 {
		debug.Log("got -exp flag, running...")
		_, err := run(strings.NewReader(*execute), "cli")
		if err != nil {
			log.Fatalln(err)
		}
	} else if len(flag.Args()) == 1 {
		debug.Log("got file, running...")
		file := flag.Args()[0]
		f, err := os.Open(file)
		if err != nil {
			log.Fatalf("Failed to open file: %s\n", err)
		}
		defer f.Close()
		_, err = run(f, file)
		if err != nil {
			log.Fatalln("\n" + err.Error())
		}
	} else {
		fmt.Print(core.ASCII_ART, "\n")
		debug.Log("got nothing, starting repl...")
		repl(run)
	}
}
