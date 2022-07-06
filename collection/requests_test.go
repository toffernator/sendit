package collection_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/toffer/sendit/collection"
)

func TestAdd(t *testing.T) {
	r1, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Log("Creating first request failed: ", err)
		t.Fail()
	}

	r2, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Log("Creating second requset failed: ", err)
		t.Fail()
	}

	t.Run("Add sequentially", func(t *testing.T) {
		collection.Add(r1)
		collection.Add(r2)
		if collection.Size() != 2 {
			t.Logf("Expected a size of 2 but is %d", collection.Size())
			t.Fail()
		}
	})

	t.Run("Add concurrently", func(t *testing.T) {
		go collection.Add(r1)
		collection.Add(r2)

		// Allow concurrent processes to complete before asserting state
		time.Sleep(10 * time.Millisecond)
		if collection.Size() != 4 {
			t.Logf("Expected a size of 4 but is %d", collection.Size())
			t.Fail()
		}
	})

	t.Cleanup(func() {
		collection.Clear()
	})
}

func TestRemove(t *testing.T) {
	r1, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Log("Creating first request failed: ", err)
		t.Fail()
	}

	r2, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Log("Creating second requset failed: ", err)
		t.Fail()
	}

	r3, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Log("Creating second requset failed: ", err)
		t.Fail()
	}

	collection.Add(r1)
	collection.Add(r2)
	collection.Add(r3)

	t.Run("Remove sequentially", func(t *testing.T) {
		collection.Remove()
		collection.Remove()
		collection.Remove()

		if collection.Size() != 0 {
			t.Logf("Expected a size of 0 but is %d", collection.Size())
			t.Fail()
		}
	})

	collection.Add(r1)
	collection.Add(r2)
	collection.Add(r3)

	t.Run("Remove concurrently", func(t *testing.T) {
		go collection.Remove()
		go collection.Remove()
		collection.Remove()

		// Allow concurrent processes to complete before asserting state
		time.Sleep(10 * time.Millisecond)
		if collection.Size() != 0 {
			t.Logf("Expected a size of 0 but is %d", collection.Size())
			t.Fail()
		}
	})

}
