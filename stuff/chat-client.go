package main

import (
	"net"
	"fmt"
	"encoding/json"
	"os"
	"bufio"
	"strings"
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

func main() {

	// connect to this socket
	conn, e := net.Dial("tcp", "127.0.0.1:5678")

	if e != nil {
		//panic("Error getting the port :5678")
		fmt.Fprintf(os.Stderr, "error: Error connecting to the port :5678\n")
		os.Exit(1)
	}

	for { 
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		mg, _ := reader.ReadString('\n')
		
		tmp := map[string]interface{}{
			"struct": "Message2",
			"message": mg,
		}

		if strings.Contains(mg,"Disconnect") {
			tmp["mode"]="Disconnect"
		} else {
			tmp["mode"]="Mg"
		}
		
		b, err := json.Marshal(tmp)
		if err != nil {
			fmt.Println("Error %s", err.Error())
		}
		
		conn.Write(b)
	}
}
