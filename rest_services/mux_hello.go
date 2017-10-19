package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func hello(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Hello World!")
}

func handleRequests(port int) {
	router := mux.newRouter().StrictSlash(false)

	router.HandleFunc("/hello", hello)

	http.ListenAndServe(":"+strconv.Itoa(port), router)
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: program <port>")
		os.Exit(1)
	}

	fmt.Println("Starting server....")

	handleRequest(strconv.Atoi(os.Args))
}
