package main
import (
	"sync"
	"fmt"
	"time"
)

var a int
var mx sync.RWMutex
var wg sync.WaitGroup

func writer (id int) {
	mx.Lock()
	a++
	aux := a
	fmt.Println("Id: ",id,", Increase: ",aux)
	mx.Unlock()

	wg.Done()
}

func reader (id int) {
	mx.RLock()
	aux := a
	fmt.Println("Id: ",id,",Read: ",aux)
	mx.RUnlock()	

	wg.Done()

}

func main() {
	wg.Add(100)
	a = 0
	for i:=0; i < 100; i++ {
		if  i%2 == 0 {
			go writer(i)
		} else {
			go reader(i)
		}
	}
	time.Sleep(100)
	wg.Wait()
}
