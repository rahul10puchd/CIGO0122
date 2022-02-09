/*
Doing work with the help of concurrency only for calling the function do()
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

//Time Consuming Task
func do() {
	defer wait.Done()
	for j := 0; j < 10000; j++ {
		for i := 0; i < 9000000; i++ {
			continue
		}
	}
}

func main() {
	start := time.Now() //to store the time when the program starts
	wait.Add(2)

	//Using GoRoutines to run the two task simultaneously (similar to the task in main.go)
	go do()
	go do()

	wait.Wait()
	fmt.Println("Done")
	fmt.Println("Time taken: ", time.Since(start)) //to calc the time taken
}
