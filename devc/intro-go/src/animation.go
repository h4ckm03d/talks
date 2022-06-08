package main

import (
	"fmt"
	"time"
)

func main() {
	const x = `.oOo`
	for i := 0; i < 100; i++ {
		fmt.Printf("\f%c", x[i%len(x)])
		time.Sleep(200 * time.Millisecond)
	}
}
