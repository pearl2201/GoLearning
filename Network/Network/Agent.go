package network

import (
	"Learn/Network/Protocol"
	"Learn/Network/common"
	"Learn/Network/core/npc"
	"fmt"
	"io"
	"net"
	"os"
)

type Agent struct {
	server       TCPServer
	conn         net.Conn
	sessionID    int
	isConnecting bool
	player       npc.NPC
}

func (agent *Agent) Start(server TCPServer, conn net.Conn, sessionID int) {
	agent.server = server
	agent.conn = conn
	agent.sessionID = sessionID

	go agent.listen()
}

func (agent *Agent) listen() {
	defer agent.conn.Close()
	defer agent.server.CloseConnection(agent.sessionID)
	defer fmt.Println("Close connection")
	agent.isConnecting = true
	for agent.isConnecting {
		b := make([]byte, Protocol.LENGTH_HEADER)

		_, err := agent.conn.Read(b[0:])

		if err != nil {
			if err == io.EOF {
				fmt.Println("eof 1, may be client close")
				return
			} else {
				fmt.Println(os.Stderr, "Fatal error: ", err.Error())
				return
			}

		}

		length := common.ByteToInt(b)

		b = make([]byte, Protocol.LENGTH_TYPE_PACKET)
		_, err = agent.conn.Read(b[0:])
		if err != nil {
			if err == io.EOF {
				fmt.Println("eof, may be client close")
				return
			} else {
				fmt.Println(os.Stderr, "Fatal error: ", err.Error())
				return
			}
		}

		pid := common.ByteToInt(b)

		buf := make([]byte, length)
		reqLen := 0
		for reqLen < length {
			tempLen, err := agent.conn.Read(buf[reqLen:])
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

		agent.resolvePackage(p)
	}
}

func (agent *Agent) resolvePackage(p Protocol.Packet) {
	switch p.PID {
	case Protocol.PK_LOGIN:
		{
			{
				parser := Protocol.PacketParser{}
				parser.DecodePacket(&p)
				msg := parser.ReadString()

				fmt.Printf("%d: %s\n", agent.sessionID, msg)
				pw := Protocol.PacketParser{}
				pw.Prepare(Protocol.PK_LOGIN_SUCCESS)
				pw.WriteInt32(agent.sessionID)
				// write user data

				agent.conn.Write(pw.Encode())

			}
		}
	case Protocol.PK_LOGOUT:
		{

		}
	case Protocol.PK_POS_PLAYER:
		{
			parser := Protocol.PacketParser{}
			parser.DecodePacket(&p)
			x := parser.ReadFloat32()
			y := parser.ReadFloat32()
			pw := Protocol.PacketParser{}
			pw.WriteInt32(agent.sessionID)
			pw.WriteFloat32(x)
			pw.WriteFloat32(y)
			pw.Prepare(Protocol.PK_POS_PLAYER)
			agent.server.broadcastChannel <- pw.Encode()

		}
	}

}
