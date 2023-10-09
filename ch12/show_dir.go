package main

import (
	"fmt"
	"io/ioutil"
	"learning-go/whg/check"
)

func main() {
	showDir("my_directory")
}

func showDir(dir string) {
	fileInfos, err := ioutil.ReadDir(dir)
	check.CheckAndLog(err)
	for _, fileInfo := range fileInfos {
		fileName := dir + "/" + fileInfo.Name()
		if fileInfo.IsDir() {
			fmt.Println("Directory:", fileName)
			showDir(fileName)
		} else {
			fmt.Println("File:", fileName)
		}
	}
}
