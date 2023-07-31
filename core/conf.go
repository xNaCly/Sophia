package core

import "log"

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
