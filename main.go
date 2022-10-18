package main

import (
	"log"
	"os"
)

func main() {
	args, err := NewArgs(os.Args)
	if err != nil {
		log.Println(err.Error())
		return
	}

	app := NewGomk()
	app.Run(args)

}
