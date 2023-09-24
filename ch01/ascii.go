package main

import "fmt"

func main() {
	for i := 0; i < 128; i++ {
		fmt.Printf("%3d\t'%c'\n", i, i)
	}
}
