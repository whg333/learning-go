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
	fmt.Printf("before %#v\n", myStruct2)

	passValue(myStruct2)
	fmt.Printf("after %#v\n", myStruct2)
}

/**
自定义的结构类型在函数参数传递方面的行为与引用类型类似，也是按值传递的，但传递的是结构的副本。
这意味着函数接收的是结构值的副本，对参数的修改不会影响原始结构的值
*/
func passValue(myStruct MyStruct) {
	fmt.Printf("passValue before %#v\n", myStruct)
	myStruct.number = 123
	myStruct.word = "test"
	myStruct.toggle = false
	fmt.Printf("passValue after %#v\n", myStruct)
}
