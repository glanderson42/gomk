package main

import (
	"log"
	"os/exec"
	"sync"
	"time"
)

func runCommand(w *sync.WaitGroup, args ...string) {
	defer w.Done()
	log.Println("go", args)
	out, err := exec.Command("go", args...).CombinedOutput()
	if len(out) != 0 {
		log.Println(string(out))
	}

	if err != nil {
		log.Println("Error at:", err.Error())
	}

	time.Sleep(100 * time.Millisecond)
}
