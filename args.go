package main

import "log"

type Args struct {
	GenerateSample   bool
	GenerateMakefile bool
	Clean            bool
	Init             bool
	ProjectName      string
}

func NewArgs(programArgs []string) (*Args, error) {
	args := new(Args)
	if checkElementInArray("--generate-sample", programArgs) {
		args.GenerateSample = true
	}

	if checkElementInArray("--generate-makefile", programArgs) {
		args.GenerateMakefile = true
	}

	if checkElementInArray("--clean", programArgs) {
		args.Clean = true
	}

	if checkElementInArray("--init", programArgs) {
		if len(programArgs) != 3 {
			log.Fatal("Argument missing")
		}
		args.Init = true
		args.ProjectName = programArgs[2]
	}

	return args, nil
}
