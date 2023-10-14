package main

import "fmt"

func main() {
	twice(sayHi)
	twice(sayBye)

	var myFunc func()
	myFunc = sayHi
	myFunc()

	var mathFunc func(int, int) float64
	mathFunc = add
	fmt.Println(mathFunc(1, 2))

}

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunc func()) {
	theFunc()
	theFunc()
}

func add(a int, b int) float64 {
	return float64(a) + float64(b)
}
