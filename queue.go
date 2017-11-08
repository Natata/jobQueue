package jq

import (
	"fmt"
	"log"
)

// JobQueue is the queue for job requests
type JobQueue struct {
	queue chan *Job
}

// NewJobQueue creates a job queue with specific size
func NewJobQueue(cap int) *JobQueue {
	return &JobQueue{
		queue: make(chan *Job, cap),
	}
}

// Enqueue enqueue a job to the job queue
// if the queue is full, it would return error
func (jq *JobQueue) Enqueue(j *Job) error {
	select {
	case jq.queue <- j:
		log.Printf("enqueue job: %v", j.id)
	default:
		return fmt.Errorf("fail to enqueue job %v, the queue is full", j.id)
	}

	return nil
}

// Dequeue dequeue a job from job queue
func (jq *JobQueue) Dequeue() *Job {
	j := <-jq.queue
	return j
}
