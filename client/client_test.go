package client_test

import (
	"testing"

	"github.com/toffer/sendit/client"
)

const (
	REQUESTS    = "requests.csv" // There are 10 requests defined in requests.csv
	BASE_TARGET = "http://localhost:7777"
)

func TestComputeResultsReturns10(t *testing.T) {
	client.ParseJobs(REQUESTS, BASE_TARGET)
	client.SendReqs()
	results := client.ComputeResults()

	if results.Total != 10 {
		t.Logf("Expected 10 results found %d", results.Total)
		t.Fail()
	} else if results.Successes != 10 {
		t.Logf("Expected 10 successes found %d", results.Successes)
		t.Fail()
	}
}
