package main

import (
	"sync"
	"fmt"
	//"time"
)

type Semaphore struct {
	n      int
	channel chan int
	rmutex *sync.Mutex
	amutex *sync.Mutex
}


func NewSemaphore (n int) (*Semaphore){

	return &Semaphore{
		n: n,
		channel: make(chan int, n),
		rmutex: &sync.Mutex{},
		amutex: &sync.Mutex{},
		
	}
}

func (s *Semaphore) Acquire(n int) {
	s.amutex.Lock()
	for i := 0; i<n; i++ {
		
		s.channel <- 1
	}
	s.n -= n
	s.amutex.Unlock()
	
}

func (s *Semaphore) Release(n int) {
	
	for i := 0; i<n; i++ {
		s.rmutex.Lock()
		<-s.channel
		s.n++
		s.rmutex.Unlock()
	}	
}

func writer (id int) {
	wg.Add(1)
	mx.Acquire(1)
	a++
	aux := a
	fmt.Println("Id: ",id,", Increase: ",aux)
	mx.Release(1)
	wg.Done()
}

var a int
var wg sync.WaitGroup
var mx = NewSemaphore(1) 

func main() {
	
	a = 0
	for i:=0; i < 100; i++ {
		go writer(i)
	}
	wg.Wait()
}
