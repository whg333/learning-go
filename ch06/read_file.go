package main

import (
	"fmt"
	"learning-go/whg/files"
)

func main() {
	total := 0.0
	numbers := files.GetFloats("data.txt")
	for i, number := range numbers {
		fmt.Printf("%d = %.2f\n", i, number)
		total += number
	}
	length := len(numbers)
	average := total / float64(length)
	fmt.Printf("average is %.2f\n", average)
}
