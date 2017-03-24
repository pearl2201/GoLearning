/*DayTimeServer*/
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	/*if len(os.Args) != 2 {
		fmt.Println("error command")
		os.Exit(2)
	}*/

	port := ":1200"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		dayTime := time.Now().String()
		conn.Write([]byte(dayTime))
		conn.Close()

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %s\n", err.Error())
		os.Exit(0)
	}

}
