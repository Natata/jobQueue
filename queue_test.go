package jq

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEnqueueDequeue(t *testing.T) {
	Convey("test job queue enqueue and dequeue", t, func(c C) {
		jq := NewJobQueue(10)

		x := 0
		expectJob := NewJob(123, func() error {
			x++
			return nil
		})

		go func() {
			err := jq.Enqueue(expectJob)
			c.So(err, ShouldBeNil)
		}()

		actualJob := jq.Dequeue()
		So(actualJob.id, ShouldEqual, expectJob.id)

		actualJob.execute()
		So(x, ShouldEqual, 1)
	})
}
