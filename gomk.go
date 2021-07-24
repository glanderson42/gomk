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
	if len(os.Args) > 2 {
		args_handler(os.Args)
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

	for _, target := range build.Targets {
		log.Println("Building:", target.Bin)

        if target.Release {
            target.Flags = append(target.Flags, `-ldflags "-s -w"`)
        }

        buildPath, _ := os.Getwd()
        log.Println(target.SourceDir)
        target.SourceDir = buildPath + string(os.PathSeparator) + target.SourceDir
		var outputPath = buildPath + string(os.PathSeparator) + target.OutputDir + string(os.PathSeparator) + target.Bin + getDefaultExtension()
		var wg sync.WaitGroup
		wg.Add(1)
		if len(target.Flags) != 0 {
            log.Println("Entering:", target.SourceDir)
            go runCommand("cd", target.SourceDir)
            log.Println("Running: go build -o", outputPath, strings.Join(target.Flags, " "), target.SourceDir)
			go runGoCommand(&wg, "build", "-o", outputPath, strings.Join(target.Flags, " "))
            log.Println("Leaving...")
            go runCommand("cd", buildPath)
		} else {
            log.Println("Entering:", target.SourceDir)
            go runCommand("cd", target.SourceDir)
            log.Println("Running: go build -o", outputPath, target.SourceDir)
			go runGoCommand(&wg, "build", "-o", outputPath)
            log.Println("Leaving...")
            go runCommand("cd", buildPath)
		}
		wg.Wait()
		log.Println("Finished:", target.Bin)
	}
}
