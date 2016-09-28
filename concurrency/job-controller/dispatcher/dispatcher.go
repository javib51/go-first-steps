/*
Starting with this code https://github.com/ivpusic/grpool
*/
package dispatcher

import "sync"

type worker struct {
	pool    chan *worker
	newTask chan Task
	stop    chan bool
}

type Task func()

func (w *worker) start() {

	go func() {

		for {
			w.pool <- w

			select {
			case job := <-w.newTask:
				job()
				

			case stop := <-w.stop:
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

func newWorker(pool chan *worker) *worker {
	return &worker{
		newTask: make(chan Task),
		stop:    make(chan bool),
		pool:    pool,
	}
}

type Dispatcher struct {
	queue chan Task
	pool  chan *worker
	stop  chan bool
	size  int
	wg    sync.WaitGroup
}

func (d *Dispatcher) Dispatch() {
	
	go func() {
		for {
			select {
			case job := <-d.queue:
				w := <-d.pool
				w.newTask <- job
			case stop := <-d.stop:
				//¿Close both channels  or use w.close()?
				//In that case this is correct
				if stop {
					for i := 0; i < d.size; i++ {
						worker := <-d.pool
						worker.stop <- true
						<-worker.stop
					}

					d.stop <- true
					return
				}
			}
		}
	}()
}

// If you are not using WaitAll
func (d *Dispatcher) Done() {
	d.wg.Done()
}

func (d *Dispatcher) Wait(n int) {
	d.wg.Add(n)
}

func (d *Dispatcher) WaitAll() {
	d.wg.Wait()
}

func (d *Dispatcher) Release() {
	d.stop <- true
	<-d.stop
}

func NewDispatcher(nWorkers int, qCapacity int) *Dispatcher {

	d := &Dispatcher{
		queue: make(chan Task, qCapacity),
		stop:  make(chan bool),
		size:  qCapacity,
		pool:  make(chan *worker, nWorkers),
	}

	for i := 0; i < nWorkers; i++ {
		w := newWorker(d.pool)
		w.start()
	}

	d.Dispatch()
	return d
}
