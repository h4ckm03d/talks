package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
)

const basePkg = "github.com/h4ckm03d/talks"
const basePathMessage = `Run 'go get %q'`

func main() {
	p, err := build.Default.Import(basePkg, "", build.FindOnly)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't find talks template files: %v\n", err)
		fmt.Fprintf(os.Stderr, basePathMessage, basePkg)
		os.Exit(1)
	}
	baseArg := fmt.Sprintf("-base=%s", p.Dir)
	args := append(os.Args[1:], baseArg)

	cmd := exec.Command("present", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
