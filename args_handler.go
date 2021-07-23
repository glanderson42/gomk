package main

import (
	"errors"
	"log"
	"os"
)

func args_handler(args []string) error {
	if args[1] == "--generate-sample" {
		log.Println("Generating sample...")
		generateSample()
		log.Println("Finished!")
		os.Exit(1)
		return nil
	}

	return errors.New("Unknown command!")
}
