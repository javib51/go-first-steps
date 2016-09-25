package main
import (
	"sync"
	"fmt"
)

var a int
var mx sync.Mutex
var wg sync.WaitGroup

func increase (id int) {

	mx.Lock()
	a++
	aux := a
	mx.Unlock()
	
	fmt.Println("Id: ",id,", Increase: ",aux)


	wg.Done()
}

func decrease (id int) {

	mx.Lock()
	a--
	aux := a
	mx.Unlock()
	fmt.Println("Id: ",id,",Decrease: ",aux)
	
	wg.Done()
}

func main() {
	wg.Add(100)
	a = 0
	for i:=0; i < 100; i++ {
		if i %2 == 0{
			go increase(i)
		} else {
			go decrease(i)
		}
	}
	wg.Wait()
}
