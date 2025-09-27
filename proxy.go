package main

import (
	"fmt"
	"io"
	"net"
)

func handleProxyConnection(client net.Conn) {
	defer client.Close()

	server, err := net.Dial("tcp", "127.0.0.1:2727")
	if err != nil {
		panic(err)
	}

	defer server.Close()

	go io.Copy(server, client)
	io.Copy(client, server)
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	fmt.Println("Proxy is running in port 6666")

	for {
		client, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleProxyConnection(client)
	}
}
