package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("connect to server error, err:", err)
		return
	}
	defer socket.Close()
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("please ender content:")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		//
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("read reply msg failed, err:", err)
			return
		}
		fmt.Println("got reply msg:", string(reply[:n]))

	}

}
