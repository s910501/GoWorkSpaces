package main

import "fmt"

func main() {
	var a interface{}
	a = 100
	switch v := a.(type) {
	case int32:
		fmt.Printf("yes :%v", v)
	}
}
