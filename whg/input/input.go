package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFloat(prompt string) (float64, error) {
	input, err := GetString(prompt)
	if err != nil {
		return 0, err
	}
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func GetString(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	//fmt.Println(input)
	input = strings.TrimSpace(input)
	return input, nil
}
