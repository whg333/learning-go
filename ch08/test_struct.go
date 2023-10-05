package main

import "fmt"

type MyStruct struct {
	number float64
	word   string
	toggle bool
}

func main() {
	var myStruct1 struct {
		number float64
		word   string
		toggle bool
	}
	myStruct1.number = 3.14
	myStruct1.word = "pie"
	myStruct1.toggle = true
	fmt.Printf("%#v\n", myStruct1)

	var myStruct2 MyStruct
	myStruct2.number = 3.14
	myStruct2.word = "pie"
	myStruct2.toggle = true
	fmt.Printf("%#v\n", myStruct2)
}
