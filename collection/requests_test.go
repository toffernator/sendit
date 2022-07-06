package bag_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/toffer/sendit/bag"
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
		bag.Add(r1)
		bag.Add(r2)
		if bag.Size() != 2 {
			t.Logf("Expected a size of 2 but is %d", bag.Size())
			t.Fail()
		}
	})

	t.Run("Add concurrently", func(t *testing.T) {
		go bag.Add(r1)
		bag.Add(r2)

		// Allow concurrent processes to complete before asserting state
		time.Sleep(10 * time.Millisecond)
		if bag.Size() != 4 {
			t.Logf("Expected a size of 4 but is %d", bag.Size())
			t.Fail()
		}
	})

	t.Cleanup(func() {
		bag.Clear()
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

	bag.Add(r1)
	bag.Add(r2)
	bag.Add(r3)

	t.Run("Remove sequentially", func(t *testing.T) {
		bag.Remove()
		bag.Remove()
		bag.Remove()

		if bag.Size() != 0 {
			t.Logf("Expected a size of 0 but is %d", bag.Size())
			t.Fail()
		}
	})

	bag.Add(r1)
	bag.Add(r2)
	bag.Add(r3)

	t.Run("Remove concurrently", func(t *testing.T) {
		go bag.Remove()
		go bag.Remove()
		bag.Remove()

		// Allow concurrent processes to complete before asserting state
		time.Sleep(10 * time.Millisecond)
		if bag.Size() != 0 {
			t.Logf("Expected a size of 0 but is %d", bag.Size())
			t.Fail()
		}
	})

}
