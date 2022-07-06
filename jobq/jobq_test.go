package jobq_test

import (
	"testing"

	"github.com/toffer/sendit/jobq"
	"github.com/toffer/sendit/models"
)

func TestEnqueue(t *testing.T) {
	job := models.NewJob("GET", "/ping", nil, 200)

	t.Run("Queue is initially empty", func(t *testing.T) {
		if !jobq.IsEmpty() {
			t.Log("Queue should start empty, it is not")
			t.Fail()
		}
	})
	t.Run("Queue is not empty after enqueuing", func(t *testing.T) {
		jobq.Enqueue(*job)

		if jobq.IsEmpty() {
			t.Log("Queue should no longer be empty, it is")
			t.Fail()
		}
	})
	t.Cleanup(func() {
		for !jobq.IsEmpty() {
			_ = jobq.Dequeue()
		}
	})
}

func TestDequeue(t *testing.T) {
	job := models.NewJob("GET", "/ping", nil, 200)
	jobq.Enqueue(*job)

	t.Run("Queue is initially empty", func(t *testing.T) {
		if jobq.IsEmpty() {
			t.Log("Queue shouldn't start empty, it is")
			t.Fail()
		}
	})
	t.Run("Dequeue makes queue empty", func(t *testing.T) {
		jobq.Dequeue()
		if !jobq.IsEmpty() {
			t.Log("Queue should be empty, it is not")
			t.Fail()
		}
	})
}
