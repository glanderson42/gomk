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
	Before    Step     `json:"before"`
	After     Step     `json:"after"`
}

func (target *Target) BuildHelper(outputPath string, buildPath string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Entering:", target.SourceDir)
	os.Chdir(target.SourceDir)
	if len(target.Flags) != 0 {
		log.Println("Running: go build -o", outputPath, strings.Join(target.Flags, " "), target.SourceDir)
		runGoCommand("build", "-o", outputPath, strings.Join(target.Flags, " "))
	} else {
		log.Println("Running: go build -o", outputPath, target.SourceDir)
		runGoCommand("build", "-o", outputPath)
	}
	log.Println("Leaving...")
	os.Chdir(buildPath)
}

func (target *Target) Build(wg *sync.WaitGroup) {
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
	target.BuildHelper(outputPath, buildPath, wg)
	log.Println("Finished:", target.Bin)
}
