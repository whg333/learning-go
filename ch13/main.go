package main

import (
	"fmt"
	"io/ioutil"
	"learning-go/whg/check"
	"net/http"
	"runtime"
	"time"
)

var ch chan int

const useGo = true

func main() {
	ch = make(chan int)

	timeFunc(runFunc, "")

	//time.Sleep(time.Second * 3)
	fmt.Println("end main()")
}

func runFunc(s string) {
	//responseSize("https://example.com")

	if useGo {
		go timeFunc(responseSize, "https://baidu.com")
		go timeFunc(responseSize, "https://example.com")
		go timeFunc(responseSize, "https://163.com")
		<-ch
		<-ch
		<-ch
	} else {
		timeFunc(responseSize, "https://baidu.com")
		timeFunc(responseSize, "https://example.com")
		timeFunc(responseSize, "https://163.com")
	}

}

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	check.CheckAndLog(err)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	check.CheckAndLog(err)

	fmt.Println(len(body))
	if useGo {
		ch <- 1
	}
}

func timeFunc(function func(s string), param string) {
	funcName := getFuncName()
	fmt.Println("执行", funcName, "开始")
	begin := time.Now()
	function(param)
	elapsed := time.Since(begin)
	fmt.Println("执行", funcName, "结束:", elapsed)
}

func getFuncName() string {
	callerFuncName := "unknown"
	needPrint := false
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			callerFuncName = fn.Name()
		}
		if needPrint {
			fmt.Printf("caller:%s, file:%s, line:%d\n", callerFuncName, file, line)
		}
	}
	return callerFuncName
}
