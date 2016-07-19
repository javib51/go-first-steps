package main

import (
	"fmt"
)


func main () {

	var channel = make (chan string)
	
	for i:=0; i<10; i++  {
		go hello(i,channel)
	}
	
	for i:=0; i<10; i++  {
		fmt.Println (<-channel)
	}

}

func hello (number int, channel chan <- string) {
	

		channel <- fmt.Println("%d:Hello: ", number)


}
