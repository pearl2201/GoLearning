package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	clientCount := 0

	allClients := make(map[net.Conn]int)

	newConnections := make(chan net.Conn)

	deadConnections := make(chan net.Conn)

	messages := make(chan string)

	server, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				fmt.Println("error: ", err)
				os.Exit(1)
			}
			newConnections <- conn
		}
	}()

	for {
		select {
		case conn := <-newConnections:
			{
				log.Printf("Accept new client, #%d\n", clientCount)
				allClients[conn] = clientCount
				clientCount++

				go func(conn net.Conn, clientId int) {
					reader := bufio.NewReader(conn)
					for {
						incoming, err := reader.ReadString('\n')
						if err != nil {
							break
						}
						messages <- fmt.Sprintf("Client %d > %s", clientId, incoming)
					}
					deadConnections <- conn
				}(conn, clientCount)
			}
		case message := <-messages:
			{
				for conn, _ := range allClients {
					go func(conn net.Conn, message string) {
						_, err := conn.Write([]byte(message))
						if err != nil {
							deadConnections <- conn
						}
					}(conn, message)
				}
			}
			log.Printf("New message: %s", message)
			log.Printf("Broadcast to %d clients", len(allClients))
		case conn := <-deadConnections:
			{
				delete(allClients, conn)
			}
		}
	}
}
