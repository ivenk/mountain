package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type command struct {
	ID  int `json:"ID"`
	Cmd int `json:"Cmd"`
}

type commandBuffer []command

var buffer commandBuffer

const PORT = 10000

var address = fmt.Sprintf(":%d", PORT)

func main() {
	fmt.Println("++ Mountain started ++")
	fmt.Printf("Running on port: %d \r\n", PORT)

	buffer = commandBuffer{}
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", commands)
	log.Fatal(http.ListenAndServe(address, nil))
}

func getAllCommands(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All commands retrieved")
	err := json.NewEncoder(w).Encode(buffer)
	if err != nil {
		fmt.Println("Error while encoding commands", err)
	}

	// clear the buffer
	buffer = commandBuffer{}
}

func addNewCommand(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New command added")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	var command command
	err := json.Unmarshal(reqBody, &command)
	if err != nil {
		log.Fatal("Error during unmarshalling!")
	}

	buffer = append(buffer, command)
}

func commands(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Instructions called")
	switch r.Method {
	case "GET":
		getAllCommands(w, r)
	case "POST":
		addNewCommand(w, r)
	}
}