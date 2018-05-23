package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//validURL checks if an url is a valid url
func validURL(url string) bool {
	return true
}
func consoleInput() { //example: ddos http://web.com 4
	newcommand := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("enter new order ([target] [count]): ")
		text, _ := newcommand.ReadString('\n')
		text = strings.TrimSpace(text)
		if text != "" {
			fmt.Println("console input command: " + text)
			t := strings.Split(text, " ")
			count, err := strconv.Atoi(t[1])
			if err != nil {
				fmt.Println("[error]: count parameter must be an integer")
				continue
			}
			if t[0] == "null" || t[0] == "none" {
				t[0] = ""
			}
			if validURL(t[0]); count >= 0 {
				currentOrder.Target = t[0]
				currentOrder.Count = count
				fmt.Println(currentOrder)
			}
		}
	}
}
