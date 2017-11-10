package main

import (
	"log"
	"os"
	"os/exec"
	"path"
)

const basePkg = "github.com/h4ckm03d/talks"
const basePathMessage = `Run 'go get %q'`

func main() {

	present := path.Join(os.Getenv("GOPATH"), "present")
	cmd := exec.Command(present, os.Args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
