package jq

// Dispatcher dispatch jobs to idol workers
type Dispatcher struct {
	workerPool chan chan *Job
	jq         *JobQueue
}

// NewDispatcher creates a dispatcher with worker number
func NewDispatcher(jq *JobQueue, wn int) *Dispatcher {
	return &Dispatcher{
		workerPool: make(chan chan *Job, wn),
		jq:         jq,
	}
}

// Run starts to dispatch job from job queue and assign to a worker
func (d *Dispatcher) Run() {
	// create and run the workers
	for i := 0; i < cap(d.workerPool); i++ {
		go func() {
			worker := NewWorker(d.workerPool)
			worker.Run()
		}()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		job := d.jq.Dequeue()
		jobChan := <-d.workerPool
		jobChan <- job
	}
}
