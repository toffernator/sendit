package jobq

import (
	"log"
	"sync"

	"github.com/toffer/sendit/models"
)

var (
	jobs []models.Job
	lock sync.Mutex
)

func init() {
	jobs = make([]models.Job, 0)
	lock = sync.Mutex{}
}

func Enqueue(j models.Job) {
	lock.Lock()
	defer lock.Unlock()

	jobs = append(jobs, j)
}

func Dequeue() models.Job {
	lock.Lock()
	defer lock.Unlock()

	if IsEmpty() {
		log.Fatalln("Cannot dequeue from an empty queue")
	}

	j := jobs[0]
	if len(jobs) > 1 {
		jobs = jobs[1:]
	} else if len(jobs) == 1 {
		jobs = make([]models.Job, 0)
	}
	return j
}

func IsEmpty() bool {
	return len(jobs) == 0
}
