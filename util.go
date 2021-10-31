package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"log"
	"os"
)

func getDefaultExtension() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}

	return ".out"
}

func generateSample() {
	sample := Build{
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

	file, _ := json.MarshalIndent(sample, "", "  ")

	_ = ioutil.WriteFile("gomk.sample.json", file, 0644)
}

func checkElementInArray(element string, arr []string) bool {
	for _, elem := range arr {
		if strings.Compare(element, elem) == 0 {
			return true
		}
	}

	return false
}

func generateMakefile(build Build) {
	var makefileString string
	makefileString += "COMPILER=go"
	var targetObject string
	phony := ".PHONY: "
	clean := "clean: \n"
	all := "all: "

	for _, elem := range build.Targets {
		makefileString += "\n"
		targetObject += elem.Bin
		targetObject += ": \n"
		targetObject += fmt.Sprintf("\tcd %s && \\\n", elem.SourceDir)
		targetObject += fmt.Sprintf("\t$(COMPILER) build -o $(shell pwd)/%s/$@ %s\n", elem.OutputDir, strings.Join(elem.Flags, " "))
		targetObject += "\n"

		phony += elem.Bin + " "
		clean += fmt.Sprintf("\trm -rf $(shell pwd)/%s/%s\n", elem.OutputDir, elem.Bin)
		all += elem.Bin + " "
	}
	makefileString += targetObject
	makefileString += "\n"
	makefileString += phony
	makefileString += "\n"
	makefileString += clean
	makefileString += "\n"
	makefileString += all

	_ = ioutil.WriteFile("Makefile", []byte(makefileString), 0644)
}

func clean(build Build) {
	for _, elem := range build.Targets {
		log.Println("Removing binary: " + elem.Bin)
		os.Remove(elem.Bin)
	}
}