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

        buildPath, _ := os.Getwd()
        target.SourceDir = buildPath + string(os.PathSeparator) + target.SourceDir
		var outputPath = buildPath + target.OutputDir + string(os.PathSeparator) + target.Bin + getDefaultExtension()
		var wg sync.WaitGroup
		wg.Add(1)
		if len(target.Flags) != 0 {
            go runCommand("cd", target.SourceDir)
			go runGoCommand(&wg, "build", "-o", outputPath, strings.Join(target.Flags, " "), target.SourceDir)
            go runCommand("cd", buildPath)
		} else {
            go runCommand("cd", target.SourceDir)
			go runGoCommand(&wg, "build", "-o", outputPath, target.SourceDir)
            go runCommand("cd", buildPath)
		}
		wg.Wait()
		log.Println("Finished:", target.Bin)
	}
}
