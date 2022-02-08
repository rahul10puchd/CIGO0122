package main

import (
	"fmt"
	"time"
)

/*
func mainj() {
	channel := make(chan string, 1)
	go func(ch chan<- string) {
		ch <- "2"
		// mess := <-ch
		fmt.Println(1)

	}(channel)
	message := <-channel
	time.Sleep(time.Second * 5)
	fmt.Println(message)
}
func main() {
	channel := make(chan string, 1)
	go func(ch chan string) {
		mess := <-ch
		//ch <- "hey from anonymous"
		// "hey from anonymous" -> ch
		fmt.Println(mess)
		fmt.Println(1)

	}(channel)
	message := "Hello from MaAIN FUNCTION"
	channel <- message
	time.Sleep(time.Second * 5)
	message = <-channel
	fmt.Println(message)
}
*/
// func main() {
// 	channel := make(chan string, 1)
// 	go func(ch chan<- string) {
// 		ch <- "2"
// 		// channel <- "3"
// 		//time.Sleep(time.Second * 5)
// 		fmt.Println(1)

// 	}(channel)
// 	message := <-channel
// 	//mess := <-channel
// 	time.Sleep(time.Second * 5)
// 	fmt.Println(message)
// 	// fmt.Println(mess)
// }

// func main() {
// 	channel := make(chan string, 1)
// 	go func() {
// 		channel <- "2"
// 		// channel <- "3"
// 		//time.Sleep(time.Second * 5)
// 		fmt.Println(1)

// 	}()
// 	message := <-channel
// 	//mess := <-channel
// 	time.Sleep(time.Second * 5)
// 	fmt.Println(message)
// 	// fmt.Println(mess)
// }

// func main() {
// 	channel := make(chan string)
// 	//var wait sync.WaitGroup
// 	//wait.Add(1)
// 	go func() {
// 		channel <- "Hello from Anonymous"
// 		//time.Sleep(time.Second * 5)
// 		fmt.Println(1)
// 		//wait.Done()
// 	}()
// 	go func() {
// 		message := <-channel
// 		// // time.Sleep(time.Second * 5)
// 		fmt.Println(message)
// 		//wait.Done()
// 	}()
// 	time.Sleep(time.Second * 5)
// 	message := <-channel
// 	// // time.Sleep(time.Second * 5)
// 	fmt.Println(message)
// 	//wait.Wait()
// }
func main() {
	f := make(chan int, 1)
	var arr [3]int
	for i := 0; i < 3; i++ {
		go firstdata(f)
		arr[i] = <-f
		time.Sleep(time.Second * 2)
	}
	for i := 0; i < 3; i++ {
		fmt.Println("arr[", i, "] = ", arr[i], "\n")
	}
	fmt.Println("final value of counter:=", c)

}

var c int

func firstdata(ch chan<- int) {
	c += 1
	ch <- c

}
