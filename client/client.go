package client

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/toffer/sendit/jobparser"
	"github.com/toffer/sendit/models"
)

var (
	client   *http.Client
	SentJobs chan models.Job
	Results  map[string]Result
)

func init() {
	client = http.DefaultClient
	SentJobs = make(chan models.Job)
	Results = make(map[string]Result)
}

// SendReqs completes all the jobs that are created by jobparser.ParseJobs()
func SendReqs() {
	wg := new(sync.WaitGroup)
	for {
		j, ok := <-jobparser.Jobs
		if !ok {
			// jobs channel is closed and drained
			log.Println("Jobs channel closed and drained")
			break
		}
		wg.Add(1)
		go sendReq(wg, j)
	}
	wg.Wait()
	close(SentJobs)
}

func sendReq(wg *sync.WaitGroup, job models.Job) {
	defer wg.Done()

	resp, err := client.Do(job.Request)
	job.Err = err
	job.Response = resp

	SentJobs <- job
}

type Result struct {
	Successes int
	Failures  int
}

func TallyResults() {
	for {
		j, ok := <-SentJobs
		if !ok {
			// Results channel is closed and drained
			log.Println("SentJobs channel closed and drained")
			break
		}

		endpoint := fmt.Sprintf("%s - %s", j.Request.URL.EscapedPath(), j.Request.Method)

		var result Result
		if r, contains := Results[endpoint]; contains {
			result = r
		} else {
			result = Result{}
		}

		if j.IsSuccessful() {
			result.Successes++
		} else {
			result.Failures++
		}

		Results[endpoint] = result
	}
}
