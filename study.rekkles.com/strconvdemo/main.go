package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "10000"
	// ret1 := int64(str)
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parse int failed err", err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1)

	// string to int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	// int to string
	a := 100
	retStr := strconv.Itoa(a)
	fmt.Printf("%#v %T\n", retStr, retStr)

	// string to bool
	boolStr := "true"
	retBool, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", retBool, retBool)

	// string to float
	floatStr := "1.234"
	retfloat, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", retfloat, retfloat)

	i := int32(97)
	ret2 := string(i)
	fmt.Println(ret2)
	ret3 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v", ret3)

}
