package main

import (
	"fmt"
	"time"
)

func amIWorking() {
	neo := make(chan string)
	neo <- "one"
	value := <-neo
	print(value)
}

func amIDefering(size int) {
	for i := 1; i <= size; i++ {
		defer fmt.Println(i)
	}
}

func amIAStuckWorkerFunc() {
	for {
		time.Sleep(10 * time.Second)
		fmt.Println("Alive")
	}
}

func main() {
	// 1
	amIDefering(1)

	// 2
	amIDefering(10)

	// 3
	amIWorking()

	// 4
	go amIAStuckWorkerFunc()

	// 5
	go amIAStuckWorkerFunc() // CANCEL ME
}
