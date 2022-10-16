package main

type Args struct {
	GenerateSample   bool
	GenerateMakefile bool
	Clean            bool
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

	return &args, nil
}
