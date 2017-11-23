package main

import "fmt"

func main() {
	hello()
}

func hello() {
	defer func() {
		fmt.Printf("defer")
	}()
	
	fmt.Println("hello")
}
