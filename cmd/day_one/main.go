package main

import "fmt"

func worker(id, sleepTimer int) {
	fmt.Printf("id: %d || sleepTimer: %d\n", id, sleepTimer)
}

func main() {

	i := 0
	for i < 10 {
		worker(i, 2)
		i++
	}
}
