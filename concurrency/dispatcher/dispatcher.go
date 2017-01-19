/*
Alternative dispatcher explained
by Rob Pike in this talk 
https://talks.golang.org/2012/waza.slide#45
*/

package dispatcher


type Task struct {
	fn func() bool
	requester chan bool
}

type worker struct {
	id int
	taskQueue chan Task
	stop chan bool
	done chan *worker
}

func newWorker(id int, done chan *worker) *worker {

	return &worker{
		id : id,
		taskQueue:make(chan Task),
		stop:make(chan bool),
		done: done,
	}
}

func (w *worker) run() {

	go func(){
		
		for {
			select {
				
			case task := <-w.taskQueue: // get Request from balancer
				task.requester <- task.fn()   // call fn and send result
				w.done <- w           // we've finished this request
			case <- w.stop:
				
				println("Worker stopped")
				close(w.taskQueue)
				close(w.stop)
				close(w.done)
				return			
			}
		}
		println("Worker ended")		
	}()

}

type Pool struct {
	taskQueue chan Task
	workers *FIFOQueue
	done chan *worker
	stop chan bool
	nworkers int
}

func newPool(taskq int, workersq int) *Pool {
	pool := &Pool{
		taskQueue: make(chan Task,taskq),
		workers: NewFIFOQueue(workersq),
		done: make(chan *worker, workersq),
		stop: make(chan bool),
		nworkers: workersq,
	}
	
	for i:= 0; i< workersq; i++{
		w := newWorker(i, pool.done)
		w.run()
		
		pool.workers.Push(w)
	}

	return pool
}

func (p *Pool) run() {

	go func(){
		var w *worker
		var req Task
		var stop bool = false
		
		for p.nworkers > 0 {
			//println("Pool loop")
			
			select {
			case req = <-p.taskQueue: // received a Request...

				//println("Pool Task")
				w = p.workers.Pop()
				for w == nil && (p.workers.Size() > 0) {
					w = p.workers.Pop()
				}
				
				if w == nil {
					
					//w <- p.done
					continue
				}
				//println("Pool loop task")
				w.taskQueue <- req
				
			case w = <-p.done: // a worker has finished ...
				if stop { // if the pool is stopping the process
					w.stop <- true
					println("Pool Worker done")
				} else {
					p.workers.Push(w)  // ...so update its info

				}
				
			case stop = <-p.stop:
				println("Pool stop")
			
				for i:=0; i< p.workers.Size(); i++ {
					w = p.workers.Pop()
					w.stop <- true
				}	
			
			}
		}
		println("Pool stopped")
		close(p.taskQueue)
		close(p.done)
		close(p.stop)
	}()
}
