package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	waitgroup := new(sync.WaitGroup)
	for jobNumber := 0; jobNumber < 10; jobNumber++ {
		waitgroup.Add(1)
		go ConcurrentJob(jobNumber, waitgroup)
	}

	waitgroup.Wait()
}

func ConcurrentJob(jobNumber int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	fmt.Printf("starting job %d\n", jobNumber)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	fmt.Printf("ending job %d\n", jobNumber)
}
