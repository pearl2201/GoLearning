package network

import (
	"fmt"
	"log"
	"net"
)

type TCPServer struct {
	listener         net.Listener
	agents           map[int](*Agent)
	exitChanel       chan int
	broadcastChannel chan []byte
	currMaxSessionID int
}

func (server *TCPServer) StartServer(port int, exitChannel chan int) {

	server.agents = make(map[int](*Agent))
	server.exitChanel = exitChannel
	server.currMaxSessionID = 0
	server.broadcastChannel = make(chan []byte)

	addr, err := net.ResolveTCPAddr("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server run on port: ", port)

	server.listener, err = net.ListenTCP("tcp", addr)
	defer server.listener.Close()

	go server.ListenBroadcastChanne()

	for {
		conn, err := server.listener.Accept()

		if err != nil {
			fmt.Println("cannot resolve client")
			continue
		}
		sessionID := server.CreateNextSessionID()
		agent := Agent{}
		server.agents[sessionID] = &agent
		go agent.Start(*server, conn, sessionID)

	}

}

func (server *TCPServer) Destroy() {
	server.listener.Close()

	for _, v := range server.agents {

		v.conn.Close()
	}
}

func (server *TCPServer) CloseConnection(sessionID int) {
	server.agents[sessionID].isConnecting = false
	server.agents[sessionID].conn.Close()
	delete(server.agents, sessionID)
}

func (server *TCPServer) CreateNextSessionID() int {
	server.currMaxSessionID++
	return server.currMaxSessionID
}

func (server *TCPServer) ListenBroadcastChanne() {
	for {
		select {
		case p := <-server.broadcastChannel:
			{
				for _, v := range server.agents {
					v.SendMessage(p)
				}
			}
		}
	}
}
