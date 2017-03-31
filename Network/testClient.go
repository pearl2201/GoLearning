package main

import (
	"Learn/Network/Protocol"
	"fmt"
	"net"
	"os"
)

func main() {
	//	messages := make(chan Protocol.Packet)
	addr, err := net.ResolveTCPAddr("tcp4", ":8080")
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: ", err.Error())
		os.Exit(2)
	}
	conn, err := net.DialTCP("tcp", nil, addr)

	writer := Protocol.PacketParser{}
	writer.Prepare(Protocol.PK_LOGIN)
	writer.WriteString("ahihi")
	msg := writer.Encode()
	for i := 0; i < 10; i++ {
		_, err = conn.Write(msg)
		fmt.Println(i)
		if err != nil {
			fmt.Println(os.Stderr, "Fatal error: ", err.Error())
			os.Exit(2)
		}
	}

	conn.Close()
}

func WriteInt(v int) []byte {
	b := make([]byte, 4)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	return b
}
