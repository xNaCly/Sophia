package debug

import (
	"log"
	"sophia/core"
)

func Log(in ...any) {
	if core.CONF.Debug {
		log.Println(in...)
	}
}
