package main

import "fmt"

func main() {
	divide(1, 0)
	fmt.Print("hello")
}
func divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered", r)
		}
	}()
	if b == 0 {
		panic("b is the issue")
	}
	return a / b
}

// func recoverB() {
// 	if r := recover(); r != nil {
// 		fmt.Println("recovered")
// 	}
// }

//if recover is not there, panic will happen, termination of program will happen
