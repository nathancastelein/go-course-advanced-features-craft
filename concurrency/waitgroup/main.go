package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for jobNumber := 0; jobNumber < 10; jobNumber++ {
		ConcurrentJob(jobNumber)
	}
}

func ConcurrentJob(jobNumber int) {
	fmt.Printf("starting job %d\n", jobNumber)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	fmt.Printf("ending job %d\n", jobNumber)
}
