package main

import (
	"fmt"
	"jobQueue"
	"sync"
	"time"
)

func main() {
	jq := jobQueue.NewJobQueue(5)
	dp := jobQueue.NewDispatcher(jq, 2)
	dp.Run()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		jid := i
		j := jobQueue.NewJob(int64(jid), func() error {
			time.Sleep(1 * time.Second)
			fmt.Printf("job %v complete!\n", jid)

			wg.Done()
			return nil
		})

		err := jq.Enqueue(j)
		if err != nil {
			fmt.Println("Enqueue fail")
			panic(err)
		}
	}

	wg.Wait()
}
