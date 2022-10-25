package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
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

	if g.project.buildParams.Before.GetCommand() != "" {
		run(g.project.buildParams.Before.GetCommand(), g.project.buildParams.Before.GetParams())
	}

	var wg sync.WaitGroup
	for _, target := range g.project.buildParams.Targets {
		if g.project.buildParams.BeforeEach.GetCommand() != "" {
			run(g.project.buildParams.BeforeEach.GetCommand(), g.project.buildParams.BeforeEach.GetParams())
		}

		if target.Before.GetCommand() != "" {
			run(target.Before.GetCommand(), target.Before.GetParams())
		}

		wg.Add(1)
		go target.Build(&wg)
		wg.Wait()

		if target.After.GetCommand() != "" {
			run(target.After.GetCommand(), target.After.GetParams())
		}

		if g.project.buildParams.AfterEach.GetCommand() != "" {
			run(g.project.buildParams.AfterEach.GetCommand(), g.project.buildParams.AfterEach.GetParams())
		}
	}

	if g.project.buildParams.After.GetCommand() != "" {
		run(g.project.buildParams.After.GetCommand(), g.project.buildParams.After.GetParams())
	}
}
