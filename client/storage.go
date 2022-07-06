package client

import (
	"sync"

	"github.com/toffer/sendit/models"
)

var (
	jobs []*models.Job
	lock sync.Mutex
)

func init() {
	jobs = make([]*models.Job, 0)
}

func addJob(j *models.Job) {
	lock.Lock()
	defer lock.Unlock()

	jobs = append(jobs, j)
}

func removeJob() *models.Job {
	lock.Lock()
	defer lock.Unlock()

	if len(jobs) > 1 {
		r := jobs[len(jobs)-1]
		jobs = jobs[:len(jobs)-1]
		return r
	} else if len(jobs) == 1 {
		r := jobs[0]
		jobs = make([]*models.Job, 0)
		return r
	} else {
		return nil
	}
}

func clearJobs() {
	lock.Lock()
	defer lock.Unlock()

	for len(jobs) > 0 {
		_ = removeJob()
	}
}
