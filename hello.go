package main

import "fmt"

//main function will be automatically called when you start the application
func main() {
	//objective - f should execute as the last statement of the main function
	// first choice is use f() at the end of the main function
	// second choice - defer
	// conn:=pg.Open()
	// defer conn.close()
	fmt.Println(1)
	fmt.Println(2)
	defer f()
	fmt.Println(3)
	defer x()
	fmt.Println(4)

}

func f() {
	fmt.Println("I am bullshit")
}
func x() {
	fmt.Println("I am shit")
}
