package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Print("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}
