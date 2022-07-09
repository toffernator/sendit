package client_test

import (
	"testing"

	"github.com/toffer/sendit/client"
	"github.com/toffer/sendit/jobparser"
)

const (
	REQUESTS    = "requests.csv" // There are 10 requests defined in requests.csv
	BASE_TARGET = "http://localhost:7777"
)

func TestComputeResultsReturns10(t *testing.T) {
	go jobparser.ParseJobs(REQUESTS, BASE_TARGET)
	go client.SendReqs()

	totalResults := 0
	for {
		_, ok := <-client.SentJobs
		if !ok {
			// results was closed and drained
			break
		}
		totalResults++
	}

	expected := 10
	if totalResults != expected {
		t.Logf("Expected %d results, but is %d", expected, totalResults)
		t.Fail()
	}
}
