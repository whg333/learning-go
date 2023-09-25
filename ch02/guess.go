package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	randomNum := getRandomNum()
	guessed := false
	for i := 10; i > 0; i-- {
		inputNum := getInputNum("[还剩余%d次]请输入你猜测的数（1到100之间）：", i)

		if inputNum < randomNum {
			fmt.Println("哎呀，你猜低了。")
		} else if inputNum > randomNum {
			fmt.Println("哎呀，你猜高了。")
		} else {
			fmt.Println("干得好！你猜对了！")
			guessed = true
			break
		}
	}

	if !guessed {
		fmt.Printf("对不起。你没有猜对。它是：[%d]", randomNum)
	}
}

func getRandomNum() int {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randomNum := rand.Intn(100) + 1
	//fmt.Println("随机生成数字为", randomNum)
	return randomNum
}

func getInputNum(prompt string, leftCount int) int {
	fmt.Printf(prompt, leftCount)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	checkAndLog(err)

	input = strings.TrimSpace(input)
	inputNum, err := strconv.Atoi(input)
	checkAndLog(err)

	//fmt.Println(inputNum)
	return inputNum
}

func checkAndLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
