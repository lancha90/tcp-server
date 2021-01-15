package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp4", ":8001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

// Handles incoming requests.
func handleConnection(conn net.Conn) {

	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	size, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	for i := 0; i < size; i++ {
		fmt.Printf("%x ", buf[i])
	}

	// Send a response back to person contacting us.
	conn.Write([]byte("Message received: "+string(buf) ) )

	// Close the connection when you're done with it.
	conn.Close()
}