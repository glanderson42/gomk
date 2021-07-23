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
	log.Println(len(os.Args))
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

		var sources string
		for index, source := range target.Sources {
			log.Println("Source:", source)
			sources += source
			if index+1 != len(target.Sources) {
				source += " "
			}
		}

		var outputPath = target.OutputDir + string(os.PathSeparator) + target.Bin + getDefaultExtension()
		var wg sync.WaitGroup
		wg.Add(1)
		if len(target.Flags) != 0 {
			go runCommand(&wg, "build", "-o", outputPath, strings.Join(target.Flags, " "), sources)
		} else {
			go runCommand(&wg, "build", "-o", outputPath, sources)
		}
		wg.Wait()
		log.Println("Finished:", target.Bin)
	}
}
