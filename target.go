package main

import (
	"log"
	"os"
	"strings"
	"sync"
)

type Target struct {
	Bin       string   `json:"bin"`
	OutputDir string   `json:"outputDir"`
	SourceDir string   `json:"sourceDir"`
	Flags     []string `json:"flags"`
	Vendor    bool     `json:"vendor"`
	Release   bool     `json:"release"`
}

func (target *Target) BuildHelper(outputPath string, buildPath string, wg *sync.WaitGroup) {
	log.Println("Entering:", target.SourceDir)
	os.Chdir(target.SourceDir)
	if len(target.Flags) != 0 {
		log.Println("Running: go build -o", outputPath, strings.Join(target.Flags, " "), target.SourceDir)
		go runGoCommand(wg, "build", "-o", outputPath, strings.Join(target.Flags, " "))
	} else {
		log.Println("Running: go build -o", outputPath, target.SourceDir)
		go runGoCommand(wg, "build", "-o", outputPath)
	}
	log.Println("Leaving...")
	os.Chdir(buildPath)
}

func (target *Target) Build() {
	log.Println("Building:", target.Bin)

	if target.Release {
		log.Println("Building mode: Release")
		target.Flags = append(target.Flags, `-ldflags -w`)
	} else {
		log.Println("Building mode: Debug")
	}

	buildPath, _ := os.Getwd()
	target.SourceDir = buildPath + string(os.PathSeparator) + target.SourceDir
	var outputPath = buildPath + string(os.PathSeparator) + target.OutputDir + string(os.PathSeparator) + target.Bin + getDefaultExtension()
	var wg sync.WaitGroup
	wg.Add(1)
	target.BuildHelper(outputPath, buildPath, &wg)
	wg.Wait()
	log.Println("Finished:", target.Bin)
}
