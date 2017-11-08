package jobQueue

import "log"

// Worker response for handle job
type Worker struct {
	workerPool chan chan *Job
	jobChan    chan *Job
}

// NewWorker creates a worker
func NewWorker(workerPool chan chan *Job) *Worker {
	return &Worker{
		workerPool: workerPool,
		jobChan:    make(chan *Job),
	}
}

// Run wait the job from dispatcher then execute it
func (w *Worker) Run() {
	for {
		w.workerPool <- w.jobChan

		select {
		case job := <-w.jobChan:
			err := job.execute()
			if err != nil {
				log.Printf("job %v execute fail: %v", job.id, err)
			}
			log.Printf("job %v execute success", job.id)
		}
	}
}
