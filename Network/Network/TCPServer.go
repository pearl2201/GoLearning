package network

import (
	"fmt"
	"log"
	"net"
)

type TCPServer struct {
	listener         net.Listener
	agents           map[int]Agent
	exitChanel       chan int
	broadcastChannel chan string
	currMaxSessionID int
}

func (server *TCPServer) startServer(port int, exitChannel chan int) {
	server.exitChanel = exitChannel
	addStr := ":" + string(port)
	addr, err := net.ResolveTCPAddr("tcp4", addStr)
	if err != nil {
		log.Fatal(err)
	}
	server.listener, err = net.ListenTCP("tcp", addr)
	defer server.listener.Close()

	go func() {
		for {
			conn, err := server.listener.Accept()
			if err != nil {
				fmt.Println("cannot resolve client")
				continue
			}
			sessionID := server.CreateNextSessionID()
			agent := Agent{}
			server.agents[sessionID] = agent
			go agent.Start(conn, sessionID)

		}
	}()

}

func (server *TCPServer) Destroy() {
	server.listener.Close()

	for _, v := range server.agents {
		v.conn.Close()
	}
}

func (server *TCPServer) CreateNextSessionID() int {
	server.currMaxSessionID++
	return server.currMaxSessionID
}
