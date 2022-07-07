package client

import (
	"net/http"
	"sync"

	jobparser "github.com/toffer/sendit/job_parser"
	"github.com/toffer/sendit/models"
)

var (
	client  *http.Client
	Results chan models.Job
)

func init() {
	client = http.DefaultClient
	Results = make(chan models.Job)
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
	close(Results)
}

type Result struct {
	Successes int
	Total     int
}

func TallyResults() Result {
	result := Result{}

	for {
		r, ok := <-Results
		if !ok {
			// Results channel is closed and drained
			break
		}

		if r.IsSuccessful() {
			result.Successes++
			result.Total++
		} else {
			result.Total++
		}
	}

	return result
}

func sendReq(wg *sync.WaitGroup, job models.Job) {
	defer wg.Done()

	resp, err := client.Do(job.Request)
	job.Err = err
	job.Response = resp

	Results <- job
}
