package main

import (
	"runtime"
	"strings"
)

func getDefaultExtension() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}

	return ""
}

func checkElementInArray(element string, arr []string) bool {
	for _, elem := range arr {
		if strings.Compare(element, elem) == 0 {
			return true
		}
	}

	return false
}
