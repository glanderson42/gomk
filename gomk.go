package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args, err := NewArgs(os.Args)
	if err != nil {
		log.Println(err.Error())
		return
	}
	args.Run()

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
		target.Build()
	}
}
