package main

import (
	"fmt"
)

type Integer int
type number int

func (i Integer) print() {
	fmt.Println(i)
}
func (i Integer) CalculateSalary() int {
	return int(i)
}

func main() {

	// j := Integer(1009)
	// j.print()
	// var a int
	// a = 10
	// a = "10"

	// fmt.Print(a)
	// pj := PermanentJob{
	// 	basicpay: 10,
	// }
	//pj.CalculateSalary()//10
	// var sc SalaryCalculator
	// sc = pj
	// total := sc.CalculateSalary()
	// fmt.Println(total)
	var ss Salary
	ss.CalculateSalary()
	fmt.Print(ss)
	// pj := PermanentJob{
	// 	basicpay: 10,
	// }
	// var s Salary
	// s = pj
	// s.CalculateSalary()
	// s.Print()
	// fmt.Println(s)
	// cj := ContractJob{
	// 	basicpay: 20,
	// }
	// fj := FreelanceJob{
	// 	basicpay: 1000,
	// }
	// i := Integer(222)

	// sc := []SalaryCalculator{pj, cj, fj, i}
	// // // total := sc.CalculateSalary()
	// // // fmt.Println(total)
	// // // sc.GetJobName()
	// totalIncome(sc)

}
func totalIncome(sc []SalaryCalculator) {
	total := 0
	for _, v := range sc {
		total = total + v.CalculateSalary()
	}
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
type PrintSalary interface {
	Print()
}
type Salary interface {
	SalaryCalculator
	PrintSalary
}

// a value needs to be captured by salary variable -
type PermanentJob struct {
	basicpay int
}

func (p PermanentJob) Print() {
	fmt.Println(p.basicpay)
}
func (p PermanentJob) CalculateSalary() int {
	return p.basicpay
}

func (p PermanentJob) GetJobName() string {
	return "Permanet Job"
}

type ContractJob struct {
	basicpay int
}

func (c ContractJob) CalculateSalary() int {
	return c.basicpay
}

type FreelanceJob struct {
	basicpay int
}

func (f FreelanceJob) CalculateSalary() int {
	return f.basicpay
}

// var sc SalaryCalculator
// sc can hold the data of multiple data types which comes with a certain condition
//sc should hold the values of types PermanentJob and ContractJob
// the condition is that types(i.e,PermanentJob and ContractJob )
//must implement all the methods of the interface
