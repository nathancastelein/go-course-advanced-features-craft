package main

import (
	"fmt"
)

func HelloWorld(done chan bool) {
	fmt.Println("Hello world from my function!")
	done <- true
}

func main() {
	done := make(chan bool)
	go HelloWorld(done)

	fmt.Println("Hello world from main!")
	<-done
}
