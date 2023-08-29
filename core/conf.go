package core

import "log"

const ASCII_ART = `
  +####
 +\    #
+  \ ß  #
+   \   # <-> ß-calculus
+ ß  \  #
 +    \#
  ++++#
`

// available targets to compile sophia to
var TARGETS = []string{
	"c",
	// "go",
	// "javascript",
	// "python",
}

type Config struct {
	Debug  bool   // enable debug logs
	Target string // target to compile sophia to
}

var CONF = Config{
	Debug: false,
}

func DbgLog(in ...any) {
	if CONF.Debug {
		log.Println(in...)
	}
}
