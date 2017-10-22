// Run mysql database as follow
// sudo docker run --name=go-mysql-basic \
// -e MYSQL_USER=gouser -e MYSQL_PASSWORD=gupw55 \
// -e MYSQL_DATABASE=hello -e MYSQL_ROOT_PASSWORD=r00t55 \
// -d mysql
// CREATE TABLE messages ( id INT NOT NULL AUTO_INCREMENT, messages TEXT NOT NULL, PRIMARY KEY(id));
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Id      int
	Message string
}

var db *sql.DB

func db_connect() *sql.DB {
	db, err := sql.Open("mysql", "gouser:gupw55@tcp(172.17.0.2)/hello")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	return db
}

func getMessages(respW http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT id, message from messages")

	if err != nil {
		fmt.Fprintf(os.Stderr, "getMessages: ERROR:\n")
	}
	defer rows.Close()

	var messages []Message

	for rows.Next() {
		var message Message
		err = rows.Scan(&message.Id, &message.Message)
		if err != nil {
			fmt.Fprintf(os.Stderr, "getMessages: ERROR:\n")
		}
		messages = append(messages, message)
	}

	json.NewEncoder(respW).Encode(messages)
}

func postMessage(respW http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var message Message
	err := decoder.Decode(&message)
	switch {
	case err == io.EOF:

	case err != nil:
		fmt.Fprintf(respW, "Request's body is incorrect\n")
	}

	defer req.Body.Close()

	stmt, _ := db.Prepare("INSERT INTO messages(message) VALUES(?)")

	stmt.Exec(message.Message)
}

func hello(respW http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(respW, "Hello World!")
}

func serve_rest_api() {
	// Create db object
	db = db_connect()

	// Create Router
	router := mux.NewRouter().StrictSlash(false)
	path := ""
	subrouter := router.PathPrefix(path).Subrouter()

	// Routes
	subrouter.HandleFunc("/hello", hello).Methods("GET")
	subrouter.HandleFunc("/messages/all", getMessages).Methods("GET")
	subrouter.HandleFunc("/messages", postMessage).Methods("POST")

	// Start serving
	http.ListenAndServe(":8080", router)
}

func main() {
	fmt.Println("Starting server...")
	serve_rest_api()
}
