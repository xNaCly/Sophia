package main

import (
	// "os"
	// "runtime/pprof"
	"sophia/core/run"
)

func main() {
	// f, err := os.Create("cpu.pprof")
	// if err != nil {
	// 	panic(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	run.Start()
}
