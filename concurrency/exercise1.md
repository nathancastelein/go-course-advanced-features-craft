# Concurrency

## Using sync.WaitGroup

Open the [WaitGroup file](./waitgroup/main.go).

Use sync.WaitGroup to run all calls to ConcurrentJob in concurrency and wait their termination.

## Using channels

Open the [Channel file](./channel/main.go).

Use a `chan bool` to run all calls to ConcurrentJob in concurrency and wait their termination.