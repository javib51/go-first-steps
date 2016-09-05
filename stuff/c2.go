package main

import (
	"fmt"
	"sync"
)

func worker (message chan string, wg *sync.WaitGroup ) {
  
	message<-"worker\n"
	wg.Done()
}

func main() {
	//wg is used to wait all the goroutines to finish
	//its a kind of mutex
	var wg sync.WaitGroup
	
	message := make(chan string,10)

	for i:=0; i < 11; i++ {
		wg.Add(1)
		go worker(message,&wg)
		
	}
	
	wg.Wait()
	close (message)
	
	for n:= range message {
		fmt.Println(n)
	}
}
