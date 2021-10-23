package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	var generateMakefile = false
	args_handler(os.Args, &generateMakefile)
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

	if generateMakefile {

	}

	for _, target := range build.Targets {
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
		if len(target.Flags) != 0 {
			log.Println("Entering:", target.SourceDir)
			os.Chdir(target.SourceDir)
			log.Println("Running: go build -o", outputPath, strings.Join(target.Flags, " "), target.SourceDir)
			go runGoCommand(&wg, "build", "-o", outputPath, strings.Join(target.Flags, " "))
			log.Println("Leaving...")
			os.Chdir(buildPath)
		} else {
			log.Println("Entering:", target.SourceDir)
			os.Chdir(target.SourceDir)
			log.Println("Running: go build -o", outputPath, target.SourceDir)
			go runGoCommand(&wg, "build", "-o", outputPath)
			log.Println("Leaving...")
			os.Chdir(buildPath)
		}
		wg.Wait()
		log.Println("Finished:", target.Bin)
	}
}
