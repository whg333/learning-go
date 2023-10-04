package main

import (
	"fmt"
	"learning-go/whg/input/args"
)

func main() {
	floats := args.GetFloats()
	average := calAverage(floats...)
	fmt.Printf("average is %.2f", average)
}

func calAverage(numbers ...float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	length := len(numbers)
	return sum / float64(length)
}
