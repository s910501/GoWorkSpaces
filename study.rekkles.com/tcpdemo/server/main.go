package main

import (
	"fmt"
	"net"
)

func processConn(conn net.Conn) {
	for {
		var tmp [128]byte
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Print("read from conn failed", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		echo := "receive" + string(tmp[:n])
		conn.Write([]byte(echo))
	}

}

func main() {
	lisener, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("Start server failed.", err)
		return
	}

	for {
		conn, err := lisener.Accept()
		if err != nil {
			fmt.Println("accept failed", err)
			return
		}
		go processConn(conn)
	}
}
