package main

import (
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer closeListener(listener)

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Connection Error", err)
	}
	handleConnection(conn)

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}

func closeListener(ln net.Listener) {
	if err := ln.Close(); err != nil {
		fmt.Println("Error while closing the listener")
	}
}

func handleConnection(conn net.Conn) {
	// Defer closing the connection gracefully
	defer func(conn net.Conn) {
		fmt.Println("Closing connection")
		err := conn.Close()
		if err != nil {
			fmt.Println("Error while closing the connection", err)
		}
	}(conn)

	// Create a 1KB buffer to read data from the connection
	buf := make([]byte, 1024)

	// Read data from the connection
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error while reading data from connection", err)
	}

	redisCommand := string(buf[:n])
	fmt.Println("Received command: ", redisCommand)

	redisResponse := processCommand(redisCommand)
	_, err = conn.Write([]byte(redisResponse))

	if err != nil {
		fmt.Println("Error while reading from the connection", err)
	}

}

func processCommand(command string) string {
	//log the command
	fmt.Println("Processing command: ", command)
	return "+PONG\r\n"
}
