package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

func main() {
	port := flag.CommandLine.Int("port", 31000, "Number of TCP port to listen on")

	flag.Parse()

	listenAddress := fmt.Sprintf(":%d", *port)
	socket, err := net.Listen("tcp", listenAddress)

	if err != nil {
		panic(err)
	}

	for {
		connection, err := socket.Accept()

		if err != nil {
			panic(err)
		}

		addr := connection.RemoteAddr().String()
		addr = strings.Split(addr, ":")[0]
		connection.Write([]byte(addr))
		connection.Close()
	}
}
