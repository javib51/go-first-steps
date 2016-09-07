/*
Starting with this code https://github.com/ivpusic/grpool
*/
package dispatcher

import (
	//"fmt"
)



type worker struct {
	pool chan *worker
	newTask chan Task
	stop chan bool	
}

type Task func()

func (w *worker) start(){

	go func(){

		for {
			//fmt.Println("Worker: 1")
			w.pool <- w

			//fmt.Println("Worker: 2")

			select {
			case job := <-w.newTask:	
				job()
				
			case stop := <- w.stop:
				//¿Close both channels  or use w.close()?
				//In that case this is correct
				if stop {
					w.stop <- true
					return
				}
			}


		}
	}()

}

func newWorker (pool chan *worker) *worker{
	return &worker {
		newTask: make(chan Task),
		stop: make(chan bool),
		pool: pool,
	}
}

type Dispatcher struct{
	queue chan Task
	pool chan *worker
	stop chan bool
}

func (d *Dispatcher) dispatch(){	

	go func(){
		for {
				//fmt.Println("Dispatcher: 3")
			select {
			case job := <-d.queue:
				w := <-d.pool
				w.newTask <- job
			case stop := <- d.stop:
				//¿Close both channels  or use w.close()?
				//In that case this is correct
				if stop {
					d.stop <- true
					return
				}
			}
			//fmt.Println("Dispatcher: 4")
		}
	}()
}

func NewDispatcher (nWorkers int, qCapacity int) *Dispatcher{

	//fmt.Println("Dispatcher: 1")
	
	d := &Dispatcher {
		queue: make(chan Task, qCapacity),
		stop: make(chan bool),
		pool: make(chan *worker, nWorkers),
	}

	for i:= 0; i< nWorkers; i++{
		w := newWorker(d.pool)
		w.start()
	}

	//fmt.Println("Dispatcher: 2")
	d.dispatch()
	//fmt.Println("Dispatcher: 22")
	return d
}



