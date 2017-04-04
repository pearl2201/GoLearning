package main

import (
	"Learn/Network/Protocol"
	"fmt"
	"io"
	"net"
	"os"
)

var (
	LENGTH_HEADER      int = 4
	LENGTH_TYPE_PACKET int = 4
)

func main() {

	messages := make(chan Protocol.Packet)

	addr, err := net.ResolveTCPAddr("tcp4", ":8080")
	checkError(err)
	listener, err := net.ListenTCP("tcp", addr)
	defer listener.Close()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("cannot resolve client")
				continue
			}
			go handleConnection(conn, messages)
		}
	}()

	for {
		select {
		case p := <-messages:
			{
				if p.PID == Protocol.PK_LOGIN {
					parser := Protocol.PacketParser{}
					parser.DecodePacket(&p)
					msg := parser.ReadString()
					fmt.Println(msg)

				}
			}
		}
	}
	exitChannel := make(chan int)
	<-exitChannel
}

func handleConnection(conn net.Conn, messages chan Protocol.Packet) {
	defer conn.Close()
	defer fmt.Println("Close connection")

	for {
		b := make([]byte, LENGTH_HEADER)

		_, err := conn.Read(b[0:])

		if err != nil {
			if err == io.EOF {
				fmt.Println("eof, may be client close")
				return
			} else {
				fmt.Println(os.Stderr, "Fatal error: ", err.Error())
				return
			}

		}

		length := byteToInt(b)

		b = make([]byte, LENGTH_TYPE_PACKET)
		_, err = conn.Read(b[0:])
		if err != nil {
			if err == io.EOF {
				fmt.Println("eof, may be client close")
				return
			} else {
				fmt.Println(os.Stderr, "Fatal error: ", err.Error())
				return
			}
		}

		pid := byteToInt(b)

		buf := make([]byte, length)
		reqLen := 0
		for reqLen < length {
			tempLen, err := conn.Read(buf[reqLen:])
			reqLen += tempLen

			if err != nil {
				if err == io.EOF && reqLen < length {
					fmt.Println(os.Stderr, "Fatal error - len: -:  ", reqLen, err.Error())
					return
				} else {
					fmt.Println(os.Stderr, "Fatal error: ", err.Error())
					return
				}

			}
		}
		p := Protocol.Packet{Data: buf, PID: pid}

		messages <- p
	}

}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: ", err.Error())
		os.Exit(2)
	}
}

func byteToInt(b []byte) int {
	return int(int(b[0]) + int(b[1])<<8 + int(b[2])<<16 + int(b[3])<<24)
}
