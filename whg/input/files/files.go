package files

import (
	"bufio"
	"learning-go/whg/check"
	"os"
	"strconv"
)

func GetFloats(fileName string) []float64 {
	var numbers []float64
	lines := GetLines(fileName)
	for _, line := range lines {
		number, err := strconv.ParseFloat(line, 64)
		check.CheckAndLog(err)
		numbers = append(numbers, number)
	}
	return numbers
}

func GetLines(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	check.CheckAndLog(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	check.CheckAndLog(err)
	check.CheckAndLog(scanner.Err())
	return lines
}
