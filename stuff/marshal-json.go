package main

import (
	"encoding/json"
	"fmt"
)

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
}
func marshal(box interface{}) string {

	// Create JSON from the instance data.
	// ... Ignore errors.
	b, _ := json.Marshal(box)
	// Convert bytes to string.
	s := "[" + string(b) + "]"
	fmt.Println("debug: ",s)
	return s
}

func unmarshal (s string) {

	// Get byte slice from string.
	bytes := []byte(s)

	// Unmarshal string into structs.
	var box []Box
	json.Unmarshal(bytes, &box)

	// Loop over structs and display them.
	for l := range box {
		fmt.Println("Width = ", box[l].Width, ", Height = ", box[l].Height, ", Color = ", box[l].Color, ", Open = ", box[l].Open) 
		//fmt.Println(box)
	}
}

func main() {
	// Create an instance of the Box struct.
	box := Box{
		Width:  10,
		Height: 20,
		Color:  "blue",
		Open:   false,
	}
	s:=marshal(box)
	unmarshal(s)
	
}
