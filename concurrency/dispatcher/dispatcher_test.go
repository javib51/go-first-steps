package dispatcher

import (
	"testing"
	"fmt"
	"runtime"
)

//initTest init two arrays
func initTest(queue int, numCPUs int) (pool *Pool) {
	
	//runtime.GOMAXPROCS(numCPUs)
	println("Using MAXPROC ", numCPUs)
	var p *Pool = newPool(queue, numCPUs)
	p.run()
	
	return p
}


func TestBasic1(t *testing.T) {

	var numCPUs int = runtime.NumCPU()
	var queue int = 10

	var pool *Pool = initTest(queue, numCPUs)
	
	var ch chan bool = make(chan bool)

	for i:=0; i<queue*10; i++ {
		fn := func() bool{
			fmt.Println("Hello ",i)
			return true
		}
		
		fmt.Println("Task")
		pool.taskQueue <- Task{fn, ch}
		fmt.Println(<-ch)
	}

	println("Test stopped")
	pool.stop <- true
}

func TestBasic2(t *testing.T) {

	
}

/*import "github.com/javib51/go-first-steps/concurrency/dispatcher"
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
*/
