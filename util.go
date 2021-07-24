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
			Target{
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

func getValueByKey(key string, arr []string) string {
	for index, element := range arr {
		if strings.Compare(key, element) == 0 && index+1 < len(arr) {
			return arr[index+1]
		}
	}

	return ""
}
