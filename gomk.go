package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Gomk struct {
	buildParams Build
}

func NewGomk(buildParams Build) *Gomk {
	gomk := new(Gomk)
	gomk.buildParams = buildParams
	return gomk
}

func (g *Gomk) Run(args *Args) {
	var shouldExit bool

	if args.GenerateSample {
		log.Println("Generating sample...")
		generateSample()
		log.Println("Finished!")
		shouldExit = true
	}

	if args.GenerateMakefile {
		log.Println("Generating Makefile...")
		generateMakefile(g.buildParams)
		shouldExit = true
	}

	if args.Clean {
		log.Println("Cleaning up...")
		clean(g.buildParams)
		shouldExit = true
	}

	if shouldExit {
		os.Exit(0)
	}

	for _, target := range g.buildParams.Targets {
		target.Build()
	}
}

func main() {
	args, err := NewArgs(os.Args)
	if err != nil {
		log.Println(err.Error())
		return
	}

	jsonFile, err := ioutil.ReadFile("gomk.json")
	if err != nil {
		log.Println(err.Error())
		return
	}

	var build Build
	err = json.Unmarshal(jsonFile, &build)

	if err != nil {
		log.Println(err.Error())
		return
	}

	app := NewGomk(build)
	app.Run(args)

}
