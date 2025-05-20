package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":11211")
	if err != nil {
		panic(err)
	}
	fmt.Println("Cache server listening on :11211")
	for {
		connection, err := ln.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		fmt.Println("connected")
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 3)

		if len(parts) == 0 {
			connection.Write([]byte("ERR Empty command\n"))
			continue
		}

		cmd := strings.ToUpper(parts[0])

		switch cmd {
		case "GET":
			//Get Cache Here
			connection.Write([]byte(fmt.Sprintf("Get Cache")))

		case "SET":
			if len(parts) != 3 {
				connection.Write([]byte("ERR Usage: SET key value\n"))
				continue
			}
			//Set Cache
			connection.Write([]byte("OK\n"))
		default:
			connection.Write([]byte("ERR Unknown command\n"))
		}
	}
}
