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

func NewGomk() *Gomk {
	gomk := new(Gomk)
	return gomk
}

func (g *Gomk) Run(args *Args) {
	if args.GenerateSample {
		log.Println("Generating sample...")
		generateSample()
		log.Println("Done!")
		return
	}

	if args.Clean {
		log.Println("Cleaning up...")
		clean(g.buildParams)
		log.Println("Done!")
		return
	}

	if args.Init {
		log.Println("Running init...")
		initProject(args.ProjectName)
		log.Println("Done!")
		return
	}

	jsonFile, err := ioutil.ReadFile("gomk.json")
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = json.Unmarshal(jsonFile, &g.buildParams)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if args.GenerateMakefile {
		log.Println("Generating Makefile...")
		generateMakefile(g.buildParams)
		return
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

	app := NewGomk()
	app.Run(args)

}
