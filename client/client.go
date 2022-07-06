package client

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/toffer/sendit/models"
)

var (
	client = http.DefaultClient
)

func ParseJobs(path string, baseTarget string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		req := parseJob(scanner.Text(), baseTarget)
		addJob(req)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error while parsing requests: %s", err)
	}
}

func parseJob(toParse string, baseTarget string) *models.Job {
	jobOpts := strings.Split(toParse, ",")
	if len(jobOpts) != 4 {
		log.Fatalf("Job '%s' has an invalid format", toParse)
	}

	method := jobOpts[0]
	url := baseTarget + jobOpts[1]
	var body io.Reader
	if jobOpts[2] == "" {
		body = nil
	} else {
		body = strings.NewReader(jobOpts[2])
	}
	var expectedStatusCode int
	if i, err := strconv.Atoi(jobOpts[3]); err == nil {
		expectedStatusCode = i
	} else {
		log.Fatalf("Failed to convert %v to int with err: %s", jobOpts[3], err)
		log.Fatalf("Found a non-integer status code %v", jobOpts[3])
	}

	req := models.NewJob(method, url, body, expectedStatusCode)
	return req
}

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
