package network

import (
	"net"
)

type Agent struct {
	conn      net.Conn
	sessionID int
}

func (agent *Agent) Start(conn net.Conn, sessionID int) {
	agent.conn = conn
}
