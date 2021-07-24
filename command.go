package main

import (
	"log"
	"os/exec"
	"sync"
	"time"
)

func runCommand(command string, args ...string) {
	out, err := exec.Command(command, args...).CombinedOutput()
    time.Sleep(100 * time.Millisecond)
	if len(out) != 0 {
		log.Println(string(out))
	}

	if err != nil {
		log.Println("Error at:", err.Error())
	}
}

func runGoCommand(w *sync.WaitGroup, args ...string) {
	defer w.Done()
    log.Println(args)
    runCommand("go", args...)
	time.Sleep(100 * time.Millisecond)
}
