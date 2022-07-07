package jobparser_test

import (
	"testing"

	jobparser "github.com/toffer/sendit/job_parser"
)

const (
	REQUESTS    = "requests.csv"
	BASE_TARGET = "http://localhost:7777"
)

func TestParseJobsGivenRequestsCSVParses10Jobs(t *testing.T) {
	go jobparser.ParseJobs(REQUESTS, BASE_TARGET)
	totalJobs := 0
	for {
		_, ok := <-jobparser.Jobs()
		if !ok {
			// jobs was closed and drained
			break
		}
		totalJobs++
	}

	expected := 10
	if totalJobs != expected {
		t.Logf("Expected %d jobs, but is %d", expected, totalJobs)
		t.Fail()
	}
}
