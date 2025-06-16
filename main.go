package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// Listen on TCP port 1080 for incoming connections
	listener, err := net.Listen("tcp", ":1080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("SOCKS5 proxy server listening on port 1080")

	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Accepted connection from", conn.RemoteAddr())

	// Read the SOCKS5 version byte
	version := make([]byte, 1)
	_, err := io.ReadFull(conn, version)
	if err != nil {
		fmt.Println("Error reading version:", err)
		return
	}

	// Check if the version is SOCKS5 (version 0x05)
	if version[0] != 0x05 {
		fmt.Println("Unsupported SOCKS version:", version[0])
		return
	}

	// Send a method response (no authentication required)
	_, err = conn.Write([]byte{0x05, 0x00})
	if err != nil {
		fmt.Println("Error writing response:", err)
		return
	}

	// Handle further SOCKS commands here (currently not implemented)
	// For now, just close the connection
	fmt.Println("Closing connection to", conn.RemoteAddr())
}
