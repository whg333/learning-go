package main

import (
	"fmt"
	"io/ioutil"
	"learning-go/whg/check"
	"path/filepath"
)

func main() {
	showDir("my_directory")
}

func showDir(path string) {
	fmt.Println(path)
	fileInfos, err := ioutil.ReadDir(path)
	check.CheckAndLog(err)
	for _, fileInfo := range fileInfos {
		filePath := filepath.Join(path, fileInfo.Name())
		if fileInfo.IsDir() {
			showDir(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}
