package main

import "fmt"

/*
	for(condition){
		statements
	}
*/
func main() {
	// for i := 0; i < 3; i++ {
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println(i)
	// 	}
	// }
	// for true {
	// 	fmt.Println("Infinite Execution")
	// }
	// i := 0
	// for i < 3 {
	// 	fmt.Println(i)
	// 	i++
	// }
	i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 3 {
			break
		}
	}
}
