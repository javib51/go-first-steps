package main

import (
	"net"
	"fmt"
	"os"
	"sync"
	"encoding/json"
	"bufio"
)


const (
	Connect int16 = 1
	Mg int16 = 2
	Disconnect int16 = 3
)

type Message2 struct {
	message string
	mode string
}


var users map[string]string
var usersWait sync.WaitGroup


func handleConnection(conn net.Conn) {
	
	fmt.Println("server: New client connected: ", conn.RemoteAddr())
	
	
	fmt.Println("debug 1")
	condition := true
	for condition{

		b, err := bufio.NewReader(conn).ReadBytes('}')
		
		var data map[string]interface{}
		err = json.Unmarshal(b, &data)
		
		//fmt.Println(data)
		
		if err != nil {
			//panic(err)
		}

	
		switch data["struct"] {

		case "Message2":
			var structure Message2
			structure.message = data["message"].(string)
			structure.mode = data["mode"].(string)
			
			switch structure.mode {
			case "Connect":			
				break
			case "Mg":
				fmt.Println(conn.RemoteAddr(),": Sent message: ",structure.message)
				break
				
			case "Disconnect":
				fmt.Println(conn.RemoteAddr(),": Disconnect")
				condition = false

			}
		}		
	}
	
	conn.Close()
}

func main () {

	users = make(map[string]string)

	ln, err := net.Listen("tcp",":5678")

	if err != nil {
		//panic("Error getting the port :5678")
		fmt.Fprintf(os.Stderr, "error: Error getting the port :5678\n")
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}
