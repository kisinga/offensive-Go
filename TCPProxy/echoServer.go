package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4242")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:4242")
	for {
		// wait for a connection
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept a connection")
		}
		go improvedEcho(conn)
	}
}
func bestEcho(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func improvedEcho(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}

	log.Printf("Read %d bytes : %s", len(s), s)

	log.Println("Writing data")

	writer := bufio.NewWriter(conn)

	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()

}
func echo(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 512)
	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("CLient disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected Error")
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		log.Println("Writing back data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}
