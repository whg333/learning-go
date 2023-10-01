package main

import (
	"fmt"
	"learning-go/whg/input"
	"log"
)

func main() {
	fahrenheit, err := input.GetFloat("Enter a temperature in Fahrenheit: ")
	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
