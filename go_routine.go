package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup
var counter int
var mutex sync.Mutex

type Counter struct {
	counter int
	mutex   sync.Mutex
}

func add(hint string) {
	for i := 0; i < 3; i++ {
		mutex.Lock()
		a := counter
		a++
		time.Sleep(time.Second * 3)
		counter = a
		fmt.Println(hint, " i:=", i, " Count:=", counter)
		mutex.Unlock()
	}
	wait.Done()
}
func main32() {
	counter = 0
	wait.Add(2)
	go add("first: ")
	go add("second: ")
	wait.Wait()
	fmt.Println("final value of counter:=", counter)
}

// func main() {
// 	var mutex sync.Mutex
// 	counter := 0
// 	for i := 1; i <= 10; i++ {
// 		go func(i int) {

// 			mutex.Lock()
// 			counter++
// 			fmt.Println(i, "=>", counter)
// 			mutex.Unlock()
// 		}(i)
// 	}
// 	time.Sleep(time.Second * 3)
// 	// fmt.Println(counter)
// }

// func main() {
// 	var wait sync.WaitGroup
// 	counter := 20000
// 	wait.Add(counter)
// 	//go hello(wait)
// 	for i := 0; i < counter; i++ {
// 		go hello(&wait, i)
// 	}
// 	defer wait.Wait()

// }
// func hello(wait *sync.WaitGroup, counter int) {
// 	fmt.Println(counter)
// 	wait.Done()
// }

// func main() {
// 	var wait sync.WaitGroup
// 	wait.Add(1)
// 	//go hello(wait)

// 	defer wait.Wait()
// 	go hello(&wait)

// }
// func hello(wait *sync.WaitGroup) {
// 	fmt.Println(1)
// 	wait.Done()
// }

// func main() {
// 	go hello()
// 	time.Sleep(time.Second * 5)
// }
// func hello() {
// 	fmt.Println(1)
// }

// func hello() {

// 	go func() {
// 		fmt.Println(10)
// 		fmt.Println(19)
// 		fmt.Println(13)
// 	}()
// }
