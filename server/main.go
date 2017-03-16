package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

var allClients map[*Client]int

type Client struct {
	outgoing   chan string
	reader     *bufio.Reader
	writer     *bufio.Writer
	conn       net.Conn
	connection *Client
}

func (client *Client) Read() {
	for {
		line, err := client.reader.ReadString('\n')
		if err == nil {
			if client.connection != nil {
				client.connection.outgoing <- line
			}
			fmt.Println("")
			fmt.Println(line)
		} else {
			break
		}

	}

	client.conn.Close()
	delete(allClients, client)
	if client.connection != nil {
		client.connection.connection = nil
	}
	client = nil
}

func (client *Client) Write() {
	for data := range client.outgoing {
		client.writer.WriteString(data)
		client.writer.Flush()
	}
}
func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}
func NewClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client{
		// incoming: make(chan string),
		outgoing: make(chan string),
		conn:     connection,
		reader:   reader,
		writer:   writer,
	}
	client.Listen()

	return client
}

func watchConsoleInput(conn net.Conn) {
	newcommand := bufio.NewReader(os.Stdin)
	for {
		text, _ := newcommand.ReadString('\n')
		text = strings.TrimSpace(text)
		// send newcommand back to client
		if text != "" {
			fmt.Println("console input command: " + text)
			fmt.Println("sending command to " + strconv.Itoa(len(allClients)) + " connected clients")
			for clientList, _ := range allClients {
				//if clientList.connection == nil {
				clientList.conn.Write([]byte(text + "\n"))
				//}
			}
			fmt.Print("enter new command: ")
		}
	}
}
func watchConnectionInput(client *Client) {
	for clientList, _ := range allClients {
		if clientList.connection == nil {
			client.connection = clientList
			clientList.connection = client
			fmt.Println("Connected")
		}

	}

	allClients[client] = 1
	fmt.Println(strconv.Itoa(len(allClients)) + " connected clients")
	fmt.Print("command to send to the clients: ")
}

func main() {
	allClients = make(map[*Client]int)
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		client := NewClient(conn)
		go watchConnectionInput(client)
		go watchConsoleInput(client.conn)
	}
}
