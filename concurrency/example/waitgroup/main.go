package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

func HelloWorld() {
	defer waitgroup.Done()
	fmt.Println("Hello world from my function!")
}

func main() {
	waitgroup.Add(1)
	go HelloWorld()

	fmt.Println("Hello world from main!")
	waitgroup.Wait()
}
