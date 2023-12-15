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

type Config struct {
	AllErrors bool
	Debug     bool // enable debug logs
}

var CONF = Config{
	Debug: false,
}
