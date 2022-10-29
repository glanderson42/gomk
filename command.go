package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func run(command string, args ...string) {
	regex, _ := regexp.Compile(`\$\{\w+\}`)
	for index, command := range args {
		if regex.MatchString(command) {
			replacer := strings.NewReplacer("${", "", "}", "")
			env_val := replacer.Replace(command)
			val := os.Getenv(env_val)
			args[index] = val
		}
	}
	out, err := exec.Command(command, args...).CombinedOutput()
	if len(out) != 0 {
		log.Print(string(out))
	}

	if err != nil {
		log.Println("Error at:", err.Error())
	}
}

func runGoCommand(args ...string) {
	run("go", args...)
}
