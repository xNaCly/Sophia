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

type Config struct {
	Debug bool // enable debug logs
}

var CONF = Config{
	Debug: false,
}

func DbgLog(in ...any) {
	if CONF.Debug {
		log.Println(in...)
	}
}
