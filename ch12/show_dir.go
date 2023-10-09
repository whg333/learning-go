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

func showDir(dir string) {
	fmt.Println(dir)
	fileInfos, err := ioutil.ReadDir(dir)
	check.CheckAndLog(err)
	for _, fileInfo := range fileInfos {
		fileName := filepath.Join(dir, fileInfo.Name())
		if fileInfo.IsDir() {
			showDir(fileName)
		} else {
			fmt.Println(fileName)
		}
	}
}
