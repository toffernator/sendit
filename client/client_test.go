package client_test

import (
	"testing"

	"github.com/toffer/sendit/client"
	"github.com/toffer/sendit/collection"
)

const (
	REQUESTS    = "requests.csv" // There are 10 requests defined in requests.csv
	BASE_TARGET = "http://localhost:7777"
)

func TestParseRequestsPopulatescollection(t *testing.T) {
	client.ParseJobs(REQUESTS, BASE_TARGET)
	if collection.SizeReqs() != 10 {
		t.Logf("Expected 10 requests but found %d", collection.SizeReqs())
		t.Fail()
	}
}
