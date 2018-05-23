package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

const PORT = "3000"

type Order struct {
	Target string `json:"target"`
	Count  int    `json:"count"`
}

var currentOrder Order

func getOrder(w http.ResponseWriter, r *http.Request) {
	jsonr, err := json.Marshal(currentOrder)
	if err != nil {
		log.Panic(err)
	}
	fmt.Fprintln(w, string(jsonr))
}

func handleRequests() {
	http.HandleFunc("/", getOrder)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func main() {
	asciiTitle := `

	              _____  _____   ____   _____
	             |  __ \|  __ \ / __ \ / ____|
	   __ _  ___ | |  | | |  | | |  | | (___
	  / _' |/ _ \| |  | | |  | | |  | |\___ \
	 | (_| | (_) | |__| | |__| | |__| |____) |
	  \__, |\___/|_____/|_____/ \____/|_____/
	   __/ |
	  |___/					v.2
	`
	color.Blue(asciiTitle)

	fmt.Println("server running at port " + PORT)

	go consoleInput()

	handleRequests()
}
