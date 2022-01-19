package main

import "fmt"

/*
1. if-else statement
	if(condition){
		//statements
	}

	if(condition){
		//statements
	}else{
		if(condition){
		//statements
		}
	}
		//statements
	}

	if(condition){
		//statements
	}else if{
		//statements
	}else{
		//statements
	}

	2. switch(data){
	case var1:
		statement
	case var2:
		statement

	default:
		statement
	}
*/
func main() {
	// fmt.Print("Enter a Number:")
	// var input int
	// fmt.Scanln(&input)

	// if input%2 == 0 {
	// 	fmt.Println("hey you are even")
	// } else {
	// 	fmt.Println("hey you are odd")
	// }
	// x := 10000
	// if x = 10; 100%2 == 0 {
	// 	fmt.Println("hey you are even")
	// 	x = 200
	// } else {
	// 	fmt.Println("hey you are odd")
	// }
	// fmt.Println(x)
	data := 100
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
	case 100:
		data = 103
		fmt.Println(data)
		fallthrough //execute the next case also
	case 11:
		data = 104
		fmt.Println(data)
		fallthrough
	case 15:
		data = 1002
		fmt.Println(data)
		fallthrough
	default:
		data = 10909
		fmt.Println(data)
	}
}
