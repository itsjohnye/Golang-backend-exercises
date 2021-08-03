//clock1 is a simple clock server, which prints the current time to the client every second.

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
			log.Fatal(err) //Fatal is a shorthand for log.Println followed by a call to os.Exit(1).
			continue
		}
		handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Fatal(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
