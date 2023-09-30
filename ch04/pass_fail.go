package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	grade, err := getFloat("Enter a grade: ")
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

func getFloat(prompt string) (float64, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	//fmt.Println(input)

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
