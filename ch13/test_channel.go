package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("执行main()开始")
	begin := time.Now()

	ch := make(chan string)
	//ch <- "hi from main"
	//<-ch
	go greeting(ch)
	fmt.Println(<-ch) // 阻塞等待从ch获取值

	elapsed := time.Since(begin)
	fmt.Println("执行main()结束:", elapsed)
}

func greeting(ch chan string) {
	fmt.Println("process greeting...")
	time.Sleep(3 * time.Second)
	ch <- "hi"
}
