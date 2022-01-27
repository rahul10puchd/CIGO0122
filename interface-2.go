package main

import "fmt"

type SalaryCalculator interface {
}

//empty interface

//describe will give us the value and type of the data passed in the i
// func describe(i interface{}) {
// 	fmt.Printf("Type= %T, value=%v \n", i, i)
// 	var num int
// 	var ok bool
// 	num, ok = i.(int)
// 	fmt.Println(num)
// 	fmt.Println(ok)

// }
func describe(i interface{}) {
	fmt.Printf("Type= %T, value=%v \n", i, i)
	switch i.(type) {
	case string:
		fmt.Println("heyy i am string and i will do ")
	case int:
		fmt.Println("heyy i am int and i will do ")
	case bool:
		fmt.Println("heyy i am bool and i will do ")
	default:
		fmt.Println("Unsuported data type")
	}

}

//embded interface

func main() {
	describe("Hello")
	describe(12)
	describe(true)
	describe(2.7)

}
