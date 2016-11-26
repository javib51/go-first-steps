/*
Alternative dispatcher explained
by Rob Pike in this talk 
https://talks.golang.org/2012/waza.slide#45
*/

package dispatcher

import (
	"sync"
	"container/heap"
)

type Task struct {
	fn func() interface{}
	requester chan interface{}
}

type worker struct {
	id int
	taskQueue chan Task
	stop chan bool
}

func newWorker(id int) *worker {
	return &worker{
		id : id,
		taskQueue:make(chan Task),
		stop:make(chan bool),
	}
}

func (w *worker) run() {
	go func(){
		var task Task
		for {
			task := <-w.requests // get Request from balancer
			task.requester <- task.fn()   // call fn and send result
			done <- w           // we've finished this request
		}
	}()
}

type Element worker
type Pool struct {
	taskQueue chan Task
	workers *FIFOQueue
	done chan *worker
	stop chan bool
	nworkers int
}

func newPool(taskq int, workersq int) *Pool {
	return &Pool{
		taskQueue: make(chan Task,taskq),
		workersQueue: FIFOQueue(workersq),
		done: make(chan worker, workersq),
		stop: make(chan bool),
	}
	var w worker
	for i:= workersq; i< workersq; i++{
		w = newWorker(i)
		w.run()
	}
}

func (p *Pool) run() {
	go func(){
		var w worker
		var req Request
		for{
			select {
			case req = <-p.taskQueue: // received a Request...
				w = p.workersQueue.Pop()
				if w == nil {
					w <- p.done
				}
				
				w.taskQueue <- req
			case w = <-p.done: // a worker has finished ...
				p.workersQueue.Push(w)  // ...so update its info
			}
		}
	}()
}
