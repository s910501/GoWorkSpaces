package main

import (
	"bufio"
	"fmt"
	"os"
)

// Get user input if have space

func useScan() {
	var s string
	fmt.Print("Please input content:")
	fmt.Scanln(&s)
	fmt.Printf("Your input content is:%s\n", s)
}

func useBufio() {
	var s string
	fmt.Print("Please input content:")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("Your input is%v", s)
}

func main() {
	// useScan()
	useBufio()
}
