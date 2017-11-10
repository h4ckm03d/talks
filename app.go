package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("present", os.Args...)
	cmd.Dir = "./slides"
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
