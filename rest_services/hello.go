package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Hello World!")
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: program <port>")
		os.Exit(1)
	}

	http.HandleFunc("/hello", hello)

	fmt.Println("Server started....")

	http.ListenAndServe(":"+os.Args[1], nil)
}
