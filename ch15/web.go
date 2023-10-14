package main

import (
	"learning-go/whg/check"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello) // 注意匹配url必须/开头
	err := http.ListenAndServe("localhost:8080", nil)
	//check.CheckAndLog(err)
	log.Fatal(err) // 不使用上面的CheckAndLog是因为ListenAndServe启动成功会一直阻塞，启动失败会立即返回err
}

func hello(response http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, Go!") // 字符串转换为byte切片
	_, err := response.Write(message)
	check.CheckAndLog(err)
}
