package main

import "fmt"

/*
Structure is just a blueprint
idea, it won't consume any of the ram space until and unless the instance of it is created

structure - instance of structure
class- object
// architecture- taj-mahal, india gate, iphone13
// 10 diffrent -consume some space

type struct_name struct{

}
*/
type Iphone13 struct {
	//attributes
	camera float32
	color  string
	ram    int
	rom    int
	price  float32
	size   float32
	old    Iphone12
	// //method
	// call func()
}
type Iphone12 struct {
	camera float32
}

//independent call function
func call(number string) {
	fmt.Println("Calling... " + number)
}

//this call method is exculsively available for the instance of iphone13 structure only
func (i *Iphone13) call(number int) { //parameters are still call by value
	i.camera = 12222 //this will be reflected otside also..because call is pointer reciever
	fmt.Println(i.camera)
}

func main() {

	d := &Iphone13{}
	d.camera = 10 //auto upgrade
	d.call(999)
	fmt.Println(d.camera)
	fmt.Println(d)

	//create instance of iphone
	//empty
	// a := Iphone13{} // an instance got created with default value for it's attributes
	// a.call(12352135)
	// fmt.Println(a)
	// a.camera = 12.4 //updating the camera attribute which priviouly was 0
	// a.color = "black"

	// //	a.speaker=12
	// b := Iphone13{}
	// b.camera = 90
	// b.call(1251245)
	// a.call(12352145)

	// //non-empty instance
	// // c := Iphone13{
	// // 	camera: 10,
	// // }
	// c := Iphone13{
	// 	camera: 10,
	// 	ram:    12,
	// 	rom:    12,
	// 	size:   6,
	// 	color:  "silver",
	// 	price:  1,
	// 	old: Iphone12{
	// 		camera: 111,
	// 	},
	// }

	// fmt.Println(c)
	var a [10]int
	fmt.Println(a)
}
