# jobQueue
simple job queue and dispatcher

## Install
go get github.com/Natata/jobQueue

## Usage

### run dispatcher
```
jq := jobQueue.NewJobQueue(5)
dp := jobQueue.NewDispatcher(jq, 2)
dp.Run()
```

### create job
```
jID := int64(5566)
job := jobQueue.NewJob(jID, func() error {
  fmt.Printf("job %v complete!\n", jid)
  return nil 
})
```

### enqueue job
```
jq.Enqueue(job)
```
