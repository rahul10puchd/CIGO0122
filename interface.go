package main

import "fmt"

func main() {
	// var a int
	// a = 10
	// a = "10"

	// fmt.Print(a)
	// pj := PermanentJob{
	// 	basicpay: 10,
	// }
	// var sc SalaryCalculator
	// sc = pj
	// total := sc.CalculateSalary()
	// fmt.Println(total)

	var sc SalaryCalculator
	sc = PermanentJob{
		basicpay: 10,
	}
	total := sc.CalculateSalary()
	fmt.Println(total)
}

//interface - data type - abstract data type
//interface have only methods signature, well we can have empty interface also
/*
type interface_name interface{
	method-1
	method-2
}
*/
//
type SalaryCalculator interface {
	CalculateSalary() int
}
type PermanentJob struct {
	basicpay int
}

func (p PermanentJob) CalculateSalary() int {
	return p.basicpay
}

// type ContractJob struct {
// 	basicpay int
// }

// var sc SalaryCalculator
// sc can hold the data of multiple data types which comes with a certain condition
//sc should hold the values of types PermanentJob and ContractJob
// the condition is that types(i.e,PermanentJob and ContractJob )
//must implement all the methods of the interface
