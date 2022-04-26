package main

import (
	"flag"
	"fmt"
)

func main() {
	// name := flag.String("name", "rekkles", "please enter name")
	// age := flag.Int("age", 9000, "please enter age")
	// married := flag.Bool("married", false, "are you married?")
	// cTime := flag.Duration("ct", time.Second, "married how long")
	// fmt.Println(*name)
	// fmt.Println(*age)
	// fmt.Println(*married)
	// fmt.Println(*cTime)
	// fmt.Printf("%T\n", *cTime)

	var name string
	var age int
	flag.StringVar(&name, "name", "rekkles", "please enter name")
	flag.IntVar(&age, "age", 18, "please enter name")
	flag.Parse() // 先解析再使用
	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(flag.Args())  // 其它参数
	fmt.Println(flag.NArg())  // 其他参数个数
	fmt.Println(flag.NFlag()) // flag 参数一个数
}

// go run .\main.go -name shenzm -age 18 -ct 10h
