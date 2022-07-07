package client

import (
	"net/http"
	"sync"

	jobparser "github.com/toffer/sendit/job_parser"
	"github.com/toffer/sendit/models"
)

var (
	client  *http.Client
	results chan models.Job
)

func init() {
	client = http.DefaultClient
	results = make(chan models.Job)
}

// SendReqs completes all the jobs that are created by jobparser.ParseJobs()
func SendReqs() {
	wg := new(sync.WaitGroup)
	for {
		j, ok := <-jobparser.Jobs()
		if !ok {
			// jobs channel is closed and drained
			break
		}
		wg.Add(1)
		go sendReq(wg, j)
	}
	wg.Wait()
	close(results)
}

func Results() chan models.Job {
	return results
}

func sendReq(wg *sync.WaitGroup, job models.Job) {
	defer wg.Done()

	resp, err := client.Do(job.Request)
	job.Err = err
	job.Response = resp

	results <- job
}
