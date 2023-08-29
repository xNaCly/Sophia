package core

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
var TARGETS = map[string]struct{}{
	"js": {},
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
