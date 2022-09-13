package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("dial failed", err)
	}

	var msg string
	// if len(os.Args) < 2 {
	// 	msg = "hello wan!"
	// } else {
	// 	msg = os.Args[1]
	// }
	for {
		fmt.Print("please enter msg")
		fmt.Scanln(&msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
		for i := 0; i < 20; i++ {
			msg := "hello hello how are you !!!"
			conn.Write([]byte(msg))
			time.Sleep(time.Second)
		}
	}

	conn.Close()
}
