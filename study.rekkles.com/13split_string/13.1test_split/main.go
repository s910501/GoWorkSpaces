package main

import (
	"fmt"

	"study.rekkles.com/13split_string/split_string"
)

func main() {
	ret := split_string.Split("babcbef", "b")
	fmt.Printf("%#v\n", ret)
	ret2 := split_string.Split("bbb", "b")
	fmt.Printf("%#v\n", ret2)
	ret3 := split_string.Split("ejosada", "b")
	fmt.Printf("%#v\n", ret3)
}
