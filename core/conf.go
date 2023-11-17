package core

const ASCII_ART = `
  ██████  ▒█████   ██▓███   ██░ ██  ██▓ ▄▄▄      
▒██    ▒ ▒██▒  ██▒▓██░  ██▒▓██░ ██▒▓██▒▒████▄    
░ ▓██▄   ▒██░  ██▒▓██░ ██▓▒▒██▀▀██░▒██▒▒██  ▀█▄  
  ▒   ██▒▒██   ██░▒██▄█▓▒ ▒░▓█ ░██ ░██░░██▄▄▄▄██ 
▒██████▒▒░ ████▓▒░▒██▒ ░  ░░▓█▒░██▓░██░ ▓█   ▓██▒
▒ ▒▓▒ ▒ ░░ ▒░▒░▒░ ▒▓▒░ ░  ░ ▒ ░░▒░▒░▓   ▒▒   ▓▒█░
░ ░▒  ░ ░  ░ ▒ ▒░ ░▒ ░      ▒ ░▒░ ░ ▒ ░  ▒   ▒▒ ░
░  ░  ░  ░ ░ ░ ▒  ░░        ░  ░░ ░ ▒ ░  ░   ▒   
      ░      ░ ░            ░  ░  ░ ░        ░  ░
`

// available targets to compile sophia to
var TARGETS = map[string]struct{}{
	"js": {},
	// "go",
	// "javascript",
	// "python",
}

type Config struct {
	Debug           bool // enable debug logs
	AllErrors       bool
	Target          string // target to compile sophia to
	EnableOptimizer bool
	Ast             bool // print ast
}

var CONF = Config{
	Debug: false,
}
