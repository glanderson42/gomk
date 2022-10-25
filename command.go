package main

import (
	"log"
	"os/exec"
)

func run(command string, args ...string) {
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
