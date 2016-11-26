package main

import "github.com/javib51/go-first-steps/concurrency/dispatcher"
func main() {
	var numCPUs int = runtime.NumCPU()
	var queue int = 100
	
	runtime.GOMAXPROCS(numCPUs)
	println("using MAXPROC ", numCPUs)
	p := NewPool(queue, numCPUs)
	p.run()
	
	for i:=0; i<queue*10; i++ {
		

	}



}
