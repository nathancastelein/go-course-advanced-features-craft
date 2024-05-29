package main

import (
	"fmt"
)

func HelloWorld() {
	fmt.Println("Hello world from my function!")
}

func main() {
	go HelloWorld()
	fmt.Println("Hello world from main!")
}
