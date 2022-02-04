package main

import (
	"fmt"
	"time"
)

func main() {
	helloCh := make(chan string, 1)
	goodByeCh := make(chan string, 1)
	quitCh := make(chan bool)
	go receiveMessage(helloCh, goodByeCh, quitCh)
	go sendMessage(helloCh, "Hello world")
	time.Sleep(time.Second)
	go sendMessage(goodByeCh, "Good Bye world")

	fmt.Println("message from quitCh=", <-quitCh)
}
func sendMessage(ch chan<- string, message string) {
	ch <- message
}
func receiveMessage(helloCh, goodByeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case message := <-helloCh:
			fmt.Println("Message from helloCh:	", message)
		case message := <-goodByeCh:
			fmt.Println("Message from goodByeCh:	", message)
		case <-time.After(time.Second * 2):
			fmt.Println("Nothing receiving in 2 seconds:	Exiting the receiveMessage goroutine")
			quitCh <- true
			break
		}
	}
}
