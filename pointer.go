package main

import "fmt"

func main() {
	// i := 10
	// j := &i
	// *j = 100
	var i int
	i = 10
	var j *int //declaration..j is empty i.e, nil
	// j := new(int) //declaration + assignment(j is not empty)
	j = &i
	*j = 100

	name := new(string)
	*name = "golang"
	fmt.Println(*name)
}
