package main

import (
	"log"
	"os"
)

type Args struct {
	GenerateSample   bool
	GenerateMakefile bool
}

func NewArgs(programArgs []string) (*Args, error) {
	var args Args
	if checkElementInArray("--generate-sample", programArgs) {
		args.GenerateSample = true
	}

	if checkElementInArray("--generate-makefile", programArgs) {
		args.GenerateMakefile = true
	}

	return &args, nil
}

func (args *Args) Run(build Build) {
	var shouldExit bool

	if args.GenerateSample {
		log.Println("Generating sample...")
		generateSample()
		log.Println("Finished!")
		shouldExit = true
	}

	if args.GenerateMakefile {
		log.Println("Generating Makefile...")
		generateMakefile(build)
		shouldExit = true
	}

	if shouldExit {
		os.Exit(0)
	}
}
