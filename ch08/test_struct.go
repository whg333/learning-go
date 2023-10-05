package main

import "fmt"

// MyStruct 使用type+struct关键字来定义结构体，类似C语言的typedef+struct
type MyStruct struct {
	number float64
	word   string
	toggle bool
}

func main() {
	// 临时定义的结构体
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
