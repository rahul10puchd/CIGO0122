/*
Doing work without using any concurrecny in the program
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() //to store the time when the program starts

	//Time Consuming Task
	for j := 0; j < 10000; j++ {
		for i := 0; i < 9000000; i++ {
			continue
		}
	}
	for j := 0; j < 10000; j++ {
		for i := 0; i < 9000000; i++ {
			continue
		}
	}
	fmt.Println("Done")
	fmt.Println("Time taken: ", time.Since(start)) //to calc the time taken
}
