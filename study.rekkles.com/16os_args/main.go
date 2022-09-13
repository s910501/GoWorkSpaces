package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("%T\n", os.Args)
}
