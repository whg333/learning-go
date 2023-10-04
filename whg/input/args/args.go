package args

import (
	"learning-go/whg/check"
	"os"
	"strconv"
)

func GetFloats() []float64 {
	var floats []float64
	strings := GetStrings()
	//fmt.Println(strings)
	for _, item := range strings {
		number, err := strconv.ParseFloat(item, 64)
		check.CheckAndLog(err)
		floats = append(floats, number)
	}
	return floats
}

func GetStrings() []string {
	var strings []string
	args := os.Args[1:]
	//fmt.Println(args)
	for _, arg := range args {
		strings = append(strings, arg)
	}
	return strings
}
