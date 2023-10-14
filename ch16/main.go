package main

import (
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
