/*
 *Simple example of peers with synchronization
 *
**/

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"net"
	//"github.com/javib51/go-first-steps/ds/"
)


//if err !=nil {
//	fmt.Fprintln(os.Stderr, "File \"",args[0],"\" doesn't exist.")
//	os.Exit(1)
//}
func check (e error) {
	if e != nil {
		panic(e)
	}
}

func main () {

	args := os.Args[1:]

	if len(args) < 3{
		fmt.Fprintln(os.Stderr, "go run main.go [file with peers] [number of peer] [port]")
		os.Exit(1)
	}
	
	f, err := os.Open(args[0])
	check(err)

	scanner := bufio.NewScanner(f)
	check(err)
	
	
	
	var peer string
	
	switch args[0]{
	case "P1":
		peer="P2"
		break
	case "P2":
		peer="P2"
		break
	case "P3":
		peer="P3"
		break
	case "P4":
		peer="P4"
		break
	}

	peers:= make(map[string]string)
	
	for scanner.Scan(){
		
		d := scanner.Text()
		
		if !strings.Contains(d,args[1]) {
			fmt.Println(strings.TrimSpace(d[0:2]),":",strings.TrimSpace(d[len(d)-15:len(d)]))
			peers[d[0:2]]=strings.TrimSpace(d[len(d)-15:len(d)])
			//fmt.Println(peers)
		} 
	}

	
	ln, err := net.Listen("tcp",":"+args[2])
	check(err)


	switch args[1]{
	case "P1":
		peer=peers["P2"]
		break
	case "P2":
		peer=peers["P3"]
		break
	case "P3":
		peer=peers["P4"]
		break
	case "P4":
		peer=peers["P1"]
		break
	}
	
	if args[1] == "P1"{
		fmt.Println(args[1],": is working in a task")
		fmt.Println(args[1],": gave control to next node ",peer)
		conn, err := net.Dial("tcp",peer)
		check(err)
		conn.Write([]byte("0"))
		conn.Close()
	}

	for count:=0;count<10;count++ {
		fmt.Println(args[1],": waiting")
		c, err := ln.Accept()
		check(err)
		c.Close()

		conn, err := net.Dial("tcp",peer)
		check(err)
		fmt.Println(args[1],": gave control to next node")
	
		conn.Write([]byte(args[1]))
		conn.Close()

	}
}
