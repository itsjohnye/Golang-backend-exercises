//netcat3 is a simple netcat clone in Go.

package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Println(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	log.Println("connection close")
	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
	}
}
