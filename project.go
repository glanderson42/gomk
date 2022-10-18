package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Project struct {
	buildParams Build
}

func NewProject(buildParams Build) *Project {
	project := new(Project)
	project.buildParams = buildParams
	return project
}

func (p *Project) Sample() Build {
	return Build{
		Project:        "sample-project",
		InstallModules: false,
		Targets: []Target{
			{
				Bin:       "sample-target",
				OutputDir: "bin",
				SourceDir: "sample-project",
				Flags: []string{
					"-race",
				},
				Vendor:  false,
				Release: false,
			},
		},
	}
}

func (p *Project) Makefile() string {
	var makefileString string
	makefileString += "COMPILER=go"
	var targetObject string
	phony := ".PHONY: "
	clean := "clean: \n"
	all := "all: "

	for _, target := range p.buildParams.Targets {
		makefileString += "\n"
		targetObject += target.Bin
		targetObject += ": \n"
		targetObject += fmt.Sprintf("\tcd %s && \\\n", target.SourceDir)
		targetObject += fmt.Sprintf("\t$(COMPILER) build -o $(shell pwd)/%s/$@ %s\n", target.OutputDir, strings.Join(target.Flags, " "))
		targetObject += "\n"

		phony += target.Bin + " "
		clean += fmt.Sprintf("\trm -rf $(shell pwd)/%s/%s\n", target.OutputDir, target.Bin)
		all += target.Bin + " "
	}
	makefileString += targetObject
	makefileString += "\n"
	makefileString += phony
	makefileString += "\n"
	makefileString += clean
	makefileString += "\n"
	makefileString += all

	return makefileString
}

func (p *Project) Init() {
	err := os.Mkdir(p.buildParams.Project, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(fmt.Sprintf("%s%s%s", p.buildParams.Project, string(os.PathSeparator), "main.go"),
		[]byte(`
		package main

		func main() {}
		`),
		fs.FileMode(os.O_CREATE))

	sample := Build{
		Project:        p.buildParams.Project,
		InstallModules: false,
		Targets: []Target{
			{
				Bin:       p.buildParams.Project,
				OutputDir: "bin",
				SourceDir: ".",
				Flags:     []string{},
				Vendor:    false,
				Release:   false,
			},
		},
	}

	file, _ := json.MarshalIndent(sample, "", "  ")

	_ = ioutil.WriteFile(fmt.Sprintf("%s%s%s", p.buildParams.Project, string(os.PathSeparator), "gomk.json"), file, fs.FileMode(os.O_CREATE))
}

func (p *Project) Clean() {
	for _, target := range p.buildParams.Targets {
		log.Printf("Removing binary: %s/%s", target.OutputDir, target.Bin)
		os.Remove(fmt.Sprintf("%s%s%s", target.OutputDir, string(os.PathSeparator), target.Bin))
	}
}

func (p *Project) GenerateSample() {
	file, _ := json.MarshalIndent(p.Sample(), "", "  ")

	_ = ioutil.WriteFile("gomk.sample.json", file, 0644)
}

func (p *Project) GenerateMakefile() {
	_ = ioutil.WriteFile("Makefile", []byte(p.Makefile()), 0644)
}
