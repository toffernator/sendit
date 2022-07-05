package bag

import (
	"net/http"
	"sync"
)

var (
	reqs []*http.Request
	lock sync.Mutex
)

func init() {
	reqs = make([]*http.Request, 0)
	lock = sync.Mutex{}
}

func Add(r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	reqs = append(reqs, r)
}

func Remove() *http.Request {
	lock.Lock()
	defer lock.Unlock()

	if len(reqs) > 1 {
		item := reqs[len(reqs)-1]
		reqs = reqs[:len(reqs)-1]
		return item
	} else if len(reqs) == 1 {
		item := reqs[0]
		reqs = make([]*http.Request, 0)
		return item
	} else {
		// There are no items left in the bag
		return nil
	}
}

func Size() int {
	return len(reqs)
}
