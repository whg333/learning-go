package main

import (
	"fmt"
	"learning-go/whg/check"
	"os"
	"text/template"
)

func main() {
	_, err := os.Stdout.Write([]byte("hello\n"))
	check.CheckAndLog(err)

	executeTemplate("Here's my template!\n", nil)
	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n", 123.4)
	executeTemplate("Dot is: {{if .}}true!{{end}}\n", true)
	executeTemplate("Dot is: {{if .}}true!{{end}}\n", false)
	executeTemplate("before: {{.}}\n{{range .}}{{.}}\n{{end}}after: {{.}}\n", []string{"1", "2", "3"})

	executeTemplate("{{.Name}}, {{.Age}}\n", Person{Name: "whg", Age: 30})

	fmt.Println(os.FileMode(0600))
	fmt.Println(os.FileMode(0755)) // 0开头的0755代表八进制，FileMode对应Unix样式的文件权限
	fmt.Println(os.FileMode(755))  // 755代表十进制
	fmt.Println(os.FileMode(0070))
}

func executeTemplate(text string, data any) {
	tmpl, err := template.New("test").Parse(text)
	check.CheckAndLog(err)
	err = tmpl.Execute(os.Stdout, data)
	check.CheckAndLog(err)
}

type Person struct {
	Name string
	Age  int
}
