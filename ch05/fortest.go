package main

import (
	"fmt"
)

func main() {
	//input.GetFloat()
	data := [3]int{1, 2, 3}
	for i, e := range data {
		fmt.Printf("%d=%d\n", i, e)
	}
}
