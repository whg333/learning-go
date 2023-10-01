package main

import (
	"fmt"
	"learning-go/whg/check"
	"learning-go/whg/input"
)

func main() {
	var total float64
	var err error
	var data [3]float64
	length := len(data)
	for i := 0; i < length; i++ {
		data[i], err = input.GetFloat("Enter a float number: ")
		if err != nil {
			check.CheckAndLog(err)
		}
		total += data[i]
	}
	for i, e := range data {
		fmt.Printf("%d = %.2f\n", i, e)
	}
	fmt.Printf("average is %.2f", total/float64(length))
}
