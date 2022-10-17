package main

type Args struct {
	GenerateSample   bool
	GenerateMakefile bool
	Clean            bool
	Init             bool
	ProjectName      string
}

func NewArgs(programArgs []string) (*Args, error) {
	var args Args
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
		args.Init = true
		args.ProjectName = programArgs[2]
	}

	return &args, nil
}
