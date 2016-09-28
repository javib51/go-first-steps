package dispatcher

import (
	//"fmt"
	"runtime"
)

func Test1() {
	//recommendated
	var numCPUs int = runtime.NumCPU()
	var queue int = 1000000
	runtime.GOMAXPROCS(numCPUs)
	println("using MAXPROC ", numCPUs)
	orchestator := NewDispatcher(numCPUs, queue)

	for i := 0; i < queue; i++ {
		orchestator.Wait(1)
		orchestator.queue <- func() {
			//fmt.Println(i, ":Im worker!")
			orchestator.Done()
		}
		
	}
	//defer orchestator.Release()
	orchestator.WaitAll()
}
