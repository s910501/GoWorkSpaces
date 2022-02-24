package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("./test.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	fmt.Printf("%T\n", fileObj)
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	fmt.Print(fileInfo.Size())
	fmt.Print(fileInfo.Name())
}
