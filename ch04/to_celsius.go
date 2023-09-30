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
	fahrenheit, err := getFloat("Enter a temperature in Fahrenheit: ")
	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
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
