package main

import (
	"fmt"
	"learning-go/whg/check"
)

func main() {
	fmt.Println("begin test")
	defer test1()
	err := test2()
	check.LogPrintln(err)
	fmt.Println("end test")
}

func test1() {
	defer fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}

func test2() error {
	defer fmt.Println("a")
	fmt.Println("b")
	return fmt.Errorf("I must quit!")
	fmt.Println("c")
	return nil
}
