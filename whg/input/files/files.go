package files

import (
	"bufio"
	"learning-go/whg/check"
	"log"
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
	file, err := os.Open(fileName) // 内部调用os.OpenFile
	defer closeFile(fileName, file)

	if os.IsNotExist(err) {
		log.Printf("[%s] Not Exist!\n", fileName)
		return nil
	}
	check.CheckAndLog(err)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	check.CheckAndLog(scanner.Err())
	return lines
}

func closeFile(fileName string, file *os.File) {
	if file == nil {
		log.Printf("[%s] is nil, not need Close!\n", fileName)
		return
	}
	log.Printf("[%s] begin Close!\n", fileName)
	err := file.Close()
	check.CheckAndLog(err)
	log.Printf("[%s] end Close!\n", fileName)
}
