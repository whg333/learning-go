package main

import (
	"fmt"
	"learning-go/whg/check"
	"learning-go/whg/input/files"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/guestbook", guestbook) // 注意匹配url必须/开头
	err := http.ListenAndServe("localhost:8080", nil)
	//check.CheckAndLog(err)
	log.Fatal(err) // 不使用上面的CheckAndLog是因为ListenAndServe启动成功会一直阻塞，启动失败会立即返回err
}

func guestbook(response http.ResponseWriter, request *http.Request) {
	signatures := files.GetLines("signatures.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("guestbook.html")
	check.CheckAndLog(err)
	err = html.Execute(response, nil)
	check.CheckAndLog(err)
}
