package main

import (
	"fmt"
	"learning-go/whg/input"
	"log"
)

func main() {
	grade, err := input.GetFloat("Enter a grade: ")
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
