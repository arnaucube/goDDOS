package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() { //server

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer ln.Close()
	defer fmt.Println("ended")

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	// accept connection on port
	c, err := ln.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
	// will listen for message to process ending in newline (\n)
	message, _ := bufio.NewReader(c).ReadString('\n')
	// output message received
	fmt.Print("client connected:", string(message), "\n")
	for {
		/*// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(c).ReadString('\n')
		// output message received
		fmt.Print("client connected:", string(message), "\n")*/

		newcommand := bufio.NewReader(os.Stdin)
		fmt.Print("command to send to the clients: ")
		text, _ := newcommand.ReadString('\n')
		// send new string back to client
		c.Write([]byte(text + "\n"))
	}
}
