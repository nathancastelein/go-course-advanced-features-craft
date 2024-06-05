package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan bool)
	for jobNumber := 0; jobNumber < 10; jobNumber++ {
		go ConcurrentJob(jobNumber, done)
	}

	for jobNumber := 0; jobNumber < 10; jobNumber++ {
		<-done
	}
}

func ConcurrentJob(jobNumber int, done chan bool) {
	defer func() {
		done <- true
	}()
	fmt.Printf("starting job %d\n", jobNumber)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	fmt.Printf("ending job %d\n", jobNumber)
}
