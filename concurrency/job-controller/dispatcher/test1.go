package dispatcher

import (
	"fmt"
	"runtime"
)

func Test1(){

	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	println("using MAXPROC ",numCPUs)
	
	var workers int = 20
	var queue int = 100
	
	orchestator := NewDispatcher(workers, queue)

	for i:=0; i < queue ; i++ {
		orchestator.queue<- func() {
			if i%2==0{
				fmt.Println(":Im worker!")
			} else {
				fmt.Println(":Im worker")
			}
		}
	}
	for {}
	
}
