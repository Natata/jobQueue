package jq

import (
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDispatchJobFromJobQueue(t *testing.T) {
	Convey("dispatch queue", t, func(c C) {
		jq := NewJobQueue(10)

		dispatcher := NewDispatcher(jq, 2)
		dispatcher.Run()

		var wg sync.WaitGroup
		wg.Add(2)

		// NOTE: for test only, should not use like it, it may have race condition
		x := 0
		job1 := NewJob(1, func() error {
			x++
			wg.Done()
			return nil
		})
		job2 := NewJob(2, func() error {
			x++
			wg.Done()
			return nil
		})

		err := jq.Enqueue(job1)
		So(err, ShouldBeNil)
		err = jq.Enqueue(job2)
		So(err, ShouldBeNil)

		wg.Wait()
		So(x, ShouldEqual, 2)

	})
}
