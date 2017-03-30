package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	messages := make(chan string)
	input := make(chan string)
	addr, err := net.ResolveTCPAddr("tcp4", ":6000")
	if err != nil {
		fmt.Println("fatal error: ", err)
		os.Exit(0)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("fatal error: ", err)
		os.Exit(0)
	}
	go handlerClient(conn, input)
	go handlerServer(conn, messages)
	for {
		select {
		case message := <-messages:
			fmt.Println(message)
		case message := <-input:
			conn.Write([]byte(message + "\n"))
		}

	}
}

func handlerClient(conn net.Conn, input chan string) {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		input <- reader.Text()
	}
}

func handlerServer(conn net.Conn, messages chan string) {
	reader := bufio.NewReader(conn)
	for {
		incoming, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		messages <- incoming
	}
}
