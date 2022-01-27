package main

import (
	"errors"
	"fmt"
)

func main() {
	// var e error
	// fmt.Println(e)
	result, err := divide(4, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}

}
func divide(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("b can't be zero")
	}
	return a / b, nil
}

type Salary struct {
	basicpay int
	err      error
}

//salary can implement the error interface methods
/// if yes, then what is the benefit of this
//defer
// panic and recovery
