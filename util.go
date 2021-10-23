package main

import (
	"encoding/json"
	"io/ioutil"
	"runtime"
	"strings"
)

func getDefaultExtension() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}

	return ""
}

func generateSample() {
	sample := Build{
		Project: "sample-project",

		Targets: []Target{
			{
				Bin:       "sample-target",
				OutputDir: "bin",
				SourceDir: "sample-project",
				Flags: []string{
					"-race",
				},
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

}
