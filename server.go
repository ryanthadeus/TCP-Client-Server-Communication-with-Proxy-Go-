package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func handleServer(conn net.Conn) {
	defer conn.Close()

	var size uint32
	err := binary.Read(conn, binary.BigEndian, &size)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, size)
	conn.Read([]byte(buffer))
	receive := string(buffer)

	fmt.Println("Server Received: " + receive)

	var response string
	if receive == "I hate netvork!" {
		response = "I hate netvork toooooooooo!"
	} else {
		response = "Your Message: " + receive
	}

	binary.Write(conn, binary.BigEndian, uint32(len(response)))
	conn.Write([]byte(response))
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:2727")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	fmt.Println("Server is running on port 2727")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleServer(conn)
	}
}
