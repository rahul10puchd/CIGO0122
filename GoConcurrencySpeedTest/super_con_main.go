/*
Doing work with the help of concurrency in the function do() and in the loops as well
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

func do() {
	var internalWait sync.WaitGroup
	defer wait.Done()
	defer internalWait.Wait()

	//Also using GoRoutines in the loop to make it faster
	go func() {
		for j := 0; j < 10000; j++ {
			for i := 0; i < 9000000; i++ {
				continue
			}
		}
		internalWait.Done()
	}()
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
