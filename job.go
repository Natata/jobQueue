package jq

// Job is the basic unit for a request
type Job struct {
	id      int64
	payload func() error
}

// NewJob creates a new job
func NewJob(id int64, payload func() error) *Job {
	return &Job{
		id:      id,
		payload: payload,
	}
}

func (j *Job) execute() error {
	return j.payload()
}
