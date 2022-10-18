package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Gomk struct {
	project Project
}

func NewGomk() *Gomk {
	gomk := new(Gomk)
	return gomk
}

func (g *Gomk) Run(args *Args) {
	if args.GenerateSample {
		log.Println("Generating sample...")
		g.project.GenerateSample()
		log.Println("Done!")
		return
	}

	if args.Clean {
		log.Println("Cleaning up...")
		g.project.Clean()
		log.Println("Done!")
		return
	}

	if args.Init {
		log.Println("Running init...")
		g.project.Init()
		log.Println("Done!")
		return
	}

	jsonFile, err := ioutil.ReadFile("gomk.json")
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = json.Unmarshal(jsonFile, &g.project.buildParams)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if args.GenerateMakefile {
		log.Println("Generating Makefile...")
		g.project.GenerateMakefile()
		return
	}

	for _, target := range g.project.buildParams.Targets {
		target.Build()
	}
}
