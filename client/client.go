package client

import (
	"log"
	"net/http"
	"sync"

	"github.com/toffer/sendit/models"
)

var (
	client = http.DefaultClient
)

func SendReqs() {
	wg := new(sync.WaitGroup)
	for _, j := range jobs {
		wg.Add(1)
		go sendReq(wg, j)
	}
	wg.Wait()
}

func sendReq(wg *sync.WaitGroup, job *models.Job) {
	defer wg.Done()

	resp, err := client.Do(job.Request)
	if err != nil {
		log.Fatal(err)
	}

	job.Response = resp
	job.HasRun = true
}

type JobsResults struct {
	Total     int
	Successes int
}

func ComputeResults() JobsResults {
	results := JobsResults{
		Total:     0,
		Successes: 0,
	}
	for _, j := range jobs {
		if j.IsSuccessful() {
			results.Successes++
			results.Total++
		} else {
			results.Total++
		}
	}
	return results
}
