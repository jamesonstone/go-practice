package main

import (
    "fmt"
    "time"
)

func worker(id int, semaphore chan struct{}) {
    // Acquire
    semaphore <- struct{}{}
    fmt.Println("Worker", id, "started")

    // Simulate work
    time.Sleep(2 * time.Second)

    fmt.Println("Worker", id, "finished")
    // Release
    <-semaphore
}

func main() {
    semaphore := make(chan struct{}, 3) // Only 3 workers can run at once

    for i := 1; i <= 10; i++ {
        go worker(i, semaphore)
    }

    // Wait for all work to finish
    time.Sleep(10 * time.Second)
}
