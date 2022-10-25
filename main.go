package main

import (
	"log"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	args, err := NewArgs(os.Args)
	if err != nil {
		log.Println(err.Error())
		return
	}

	app := NewGomk()
	app.Run(args)

}
