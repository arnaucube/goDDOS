package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() { //client https://systembash.com/a-simple-go-tcp-server-and-tcp-client/

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:3333")
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("connecting to server as hostname:", hostname)
	ip := getIp()
	fmt.Println("local ip:", ip)
	// send to socket
	fmt.Fprintf(conn, hostname+", ip: "+ip.String()+"\n")
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Println("server disconnected")
			os.Exit(1)
		}
		fmt.Print("Command from server: " + message)
	}

}

func getIp() net.IP {
	var ip net.IP
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		for _, addr := range addrs {

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
		}
	}
	return ip
}
