package main

import (
	"fmt"
	"time"
)

func main3() {

	// this function returns the present time
	current_time := time.Now()

	// using inbuilt format constants
	// shown in the table above
	fmt.Println("ANSIC: ", current_time.Format(time.ANSIC))
	fmt.Println("UnixDate: ", current_time.Format(time.UnixDate))
	fmt.Println("RFC1123: ", current_time.Format(time.RFC1123))
	fmt.Println("RFC3339Nano: ", current_time.Format(time.RFC3339Nano))
	fmt.Println("RubyDate: ", current_time.Format(time.RubyDate))
}
