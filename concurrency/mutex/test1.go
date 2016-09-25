package main
import (
	"sync"
	"fmt"
)

var a int
var mx sync.Mutex
var wg sync.WaitGroup

func operation (id int) {

	mx.Lock()
	aux := a
	fmt.Println("Id: ",id,",Op: ",aux)
	a++
	mx.Unlock()
	wg.Done()
}

func main() {
	wg.Add(100)
	a = 0
	for i:=0; i < 100; i++ {
		go operation(i)
	
	}
	wg.Wait()
}
