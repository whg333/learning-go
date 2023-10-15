package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("begin main", time.Now())
	ch := make(chan string, 3)
	go sendLetters(ch)
	time.Sleep(5 * time.Second)
	fmt.Println(<-ch, time.Now())
	fmt.Println(<-ch, time.Now())
	fmt.Println(<-ch, time.Now())
	fmt.Println(<-ch, time.Now())
	fmt.Println("end main", time.Now())
}

func sendLetters(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "a"
	time.Sleep(1 * time.Second)
	ch <- "b"
	time.Sleep(1 * time.Second)
	ch <- "c"
	time.Sleep(1 * time.Second)
	ch <- "d"
}
