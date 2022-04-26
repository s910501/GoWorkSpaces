package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen UDP failed, err:", err)
		return
	}
	// 先处理错误再DEFER
	defer conn.Close()
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from UDP failed, err:", err)
			return
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		// Send data
		conn.WriteToUDP([]byte(reply), addr)
	}
}
