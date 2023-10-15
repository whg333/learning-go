package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	english := "ABCDE"
	chinese := "测试下中文"
	fmt.Println(english)
	fmt.Println(chinese)

	fmt.Println(len(english))
	fmt.Println(len(chinese))
	fmt.Println(utf8.RuneCountInString(english))
	fmt.Println(utf8.RuneCountInString(chinese))

	englishBytes := []byte(english)
	chineseBytes := []byte(chinese)
	englishBytesPart := englishBytes[2:]
	chineseBytesPart := chineseBytes[2:]
	fmt.Println(string(englishBytesPart))
	fmt.Println(string(chineseBytesPart))

	englishRunes := []rune(english)
	chineseRunes := []rune(chinese)
	englishRunesPart := englishRunes[2:]
	chineseRunesPart := chineseRunes[2:]
	fmt.Println(string(englishRunesPart))
	fmt.Println(string(chineseRunesPart))

	for index, curByte := range englishBytes {
		fmt.Printf("%d: %s\n", index, string(curByte))
	}
	for index, curByte := range chineseBytes {
		fmt.Printf("%d: %s\n", index, string(curByte))
	}

	for position, curRune := range english {
		fmt.Printf("%d: %s\n", position, string(curRune))
	}
	for position, curRune := range chinese {
		fmt.Printf("%d: %s\n", position, string(curRune))
	}
}
