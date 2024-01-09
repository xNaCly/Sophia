package run

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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
	// INFO: check if stdin is readable and the process is in a pipe
	if err == nil && !(stdinInf.Mode()&os.ModeNamedPipe == 0) {
		debug.Log("got stdin content, running...")
		out, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalln("failed to read from stdin", err)
		}
		_, err = run(string(out), "stdin")
		if err != nil {
			log.Fatalln(err)
		}
	} else if len(*execute) != 0 {
		debug.Log("got -exp flag, running...")
		_, err := run(*execute, "cli")
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
		_, err = run(string(f), file)
		if err != nil {
			log.Fatalln("\n" + err.Error())
		}
	} else {
		fmt.Print(core.ASCII_ART, "\n")
		debug.Log("got nothing, starting repl...")
		repl(run)
	}
}
