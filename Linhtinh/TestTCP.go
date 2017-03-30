package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", ":8080")
	checkError(err)
	listener, err := net.ListenTCP("tcp", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("cannot resolve client")
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}

		}

	}
	s := result.Bytes()
	fmt.Println(string(s[:]))
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: ", err.Error())
		os.Exit(2)
	}
}
