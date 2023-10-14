package main

import (
	"fmt"
	"learning-go/whg/check"
	"learning-go/whg/input/files"
	"log"
	"net/http"
	"os"
	"text/template"
)

const signatureFile = "signatures.txt"

func main() {
	http.HandleFunc("/guestbook", guestbook) // 注意匹配url必须/开头
	http.HandleFunc("/guestbook/add", add)
	http.HandleFunc("/guestbook/create", create)

	server := &http.Server{Addr: "localhost:8080", Handler: nil}
	log.Printf("Server[%s] Starting...", server.Addr)
	err := server.ListenAndServe()
	//check.CheckAndLog(err)
	log.Fatal(err) // 不使用上面的CheckAndLog是因为ListenAndServe启动成功会一直阻塞，启动失败会立即返回err
}

func guestbook(response http.ResponseWriter, request *http.Request) {
	signatures := files.GetLines(signatureFile)
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("guestbook.html")
	check.CheckAndLog(err)
	err = html.Execute(response,
		Guestbook{SignatureCount: len(signatures), Signatures: signatures})
	check.CheckAndLog(err)
}

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func add(response http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("add.html")
	check.CheckAndLog(err)
	err = html.Execute(response, nil)
	check.CheckAndLog(err)
}

func create(response http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile(signatureFile, options, os.FileMode(0600))
	check.CheckAndLog(err)
	_, err = fmt.Fprintln(file, signature)
	check.CheckAndLog(err)
	err = file.Close()
	check.CheckAndLog(err)

	http.Redirect(response, request, "/guestbook", http.StatusFound)
}
