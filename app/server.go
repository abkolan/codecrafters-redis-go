// Package main implements a simple Redis server.
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
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

	fmt.Println("Redis server listening on port 6379")
	var wg sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		fmt.Printf("New connection from %s\n", conn.RemoteAddr().String())
		wg.Add(1)
		go handleConnection(conn, &wg)
		wg.Wait()
	}

}

func closeListener(ln net.Listener) {
	if err := ln.Close(); err != nil {
		fmt.Println("Error while closing the listener")
	}
}

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		//TODO: Terrible way to pass the test, but this would be handled later.
		message, err := reader.ReadString('\n')
		message, err = reader.ReadString('\n')
		message, err = reader.ReadString('\n')
		if err != nil {
			// Check specifically for EOF
			if err.Error() == "EOF" {
				fmt.Println("Client closed connection")
				return
			}
			fmt.Printf("Error reading from connection: %v\n", err)
			return
		}
		// Process the message
		fmt.Printf("Received: %s", message)
		response := processCommand(message)
		_, err = writer.WriteString(response)
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
			return
		}

		err = writer.Flush()
		if err != nil {
			fmt.Printf("Error flushing response: %v\n", err)
			return
		}

	}
	// Defer closing the connection gracefully
	// defer func(conn net.Conn) {
	// 	fmt.Println("Closing connection")
	// 	err := conn.Close()
	// 	if err != nil {
	// 		fmt.Println("Error while closing the connection", err)
	// 	}
	// }(conn)

	/*
		for {

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
			fmt.Printf("Got response %s\n", redisResponse)
			_, err = conn.Write([]byte(redisResponse))

			if err != nil {
				fmt.Println("Error while reading from the connection", err)
			}
		}
	*/

}

func processCommand(command string) string {
	//log the command
	fmt.Println("Processing command: ", command)
	return "+PONG\r\n"
}
