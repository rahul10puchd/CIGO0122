package main

import (
	"fmt"
	"sync"
)

var arr [3]int //Package-Level Variable

func addFirstElement(wait *sync.WaitGroup, mutex *sync.Mutex) {
	defer wait.Done()
	mutex.Lock()
	//\/\/\/\/\/\/\
	arr[0] = 1 //\/
	//\/\/\/\/\/\/\
	mutex.Unlock()
}

func addSecondElement(wait *sync.WaitGroup, mutex *sync.Mutex) {
	defer wait.Done()
	mutex.Lock()
	//\/\/\/\/\/\/\
	arr[1] = 2 //\/
	//\/\/\/\/\/\/\
	mutex.Unlock()
}

func addThirdElement(wait *sync.WaitGroup, mutex *sync.Mutex) {
	defer wait.Done()
	mutex.Lock()
	//\/\/\/\/\/\/\
	arr[2] = 3 //\/
	//\/\/\/\/\/\/\
	mutex.Unlock()
}

func valueInserter(channel chan [3]int) {
	var wait sync.WaitGroup
	var mutex sync.Mutex
	wait.Add(3) //For three go routines

	//GoRoutines
	go addFirstElement(&wait, &mutex)
	go addSecondElement(&wait, &mutex)
	go addThirdElement(&wait, &mutex)

	//Waiting for all the go routines to get executed
	wait.Wait()

	//Sending data to main via a channel
	channel <- arr

}

func main() {
	channel := make(chan [3]int, 1)
	valueInserter(channel)
	assignedArray := <-channel

	//Output
	fmt.Println(assignedArray)
}
