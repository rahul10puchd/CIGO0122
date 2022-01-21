package main

import "fmt"

/*
func function_name(){
	statement-1
	statement-2
	statement-3
	statement-4
}
*/
func main() {
	w := new(int)
	//name := new(string)
	t := "hjbjk"
	_, x := c(*w, &t)
	// fmt.Println(result)
	fmt.Println(x)
	// fmt.Println(*name)
	fmt.Println(*w)
	// r, _ := b(1, 2, 3, 4, 5, 6, 6)
	// fmt.Print(r)
}
func a() (int, string) {
	return 122, "weqrr3"
}
func b(args ...int) (bool, int) {
	for _, v := range args {
		fmt.Println(v)
	}
	return true, 11
}
func c(w int, name *string) (i int, j string) {
	i = 10
	j = "rahul"
	w = 100
	*name = "code"
	return
}
