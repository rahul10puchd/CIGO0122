package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	counter := 20000
	wait.Add(counter)
	//go hello(wait)
	for i := 0; i < counter; i++ {
		go hello(&wait, i)
	}
	defer wait.Wait()

}
func hello(wait *sync.WaitGroup, counter int) {
	fmt.Println(counter)
	wait.Done()
}

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
