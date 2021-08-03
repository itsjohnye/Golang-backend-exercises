//clock2 is a simple clock server, supporting multiple clients.

package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //not fatal
			continue
		}
		go handleConn(conn) //concurrent
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Print(err) //not fatal
			return
		}
		time.Sleep(1 * time.Second)
	}
}
