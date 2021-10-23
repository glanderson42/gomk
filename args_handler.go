package main

import (
	"errors"
	"log"
	"os"
)

func args_handler(args []string, generateMakefile *bool) error {
	if checkElementInArray("--generate-sample", args) {
		log.Println("Generating sample...")
		generateSample()
		log.Println("Finished!")
		os.Exit(0)
		return nil
	} else if checkElementInArray("--generate-makefile", args) {
		log.Println("Generating makefile...")
		os.Exit(0)
		return nil
	}

	return errors.New("unknown command")
}
