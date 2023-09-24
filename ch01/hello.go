package main

import "fmt"

func main() {
	a := 4
	a = 5
	length, width := 1.2, 2.4
	var name string = "Damon Cole"
	fmt.Println("has ordered", a, "sheets")
	fmt.Println(length*width, "square meters")
	fmt.Println(name)

	Ab()
}

func Ab() {
	fmt.Println("call ab")
}
