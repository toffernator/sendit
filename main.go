package main

import (
	"fmt"

	"github.com/toffer/sendit/client"
)

const (
	REQUESTS = ".local/requests.csv"
	TARGET   = "http://localhost:8080"
)

func main() {
	client.ParseJobs(REQUESTS, TARGET)
	client.SendReqs()
	results := client.ComputeResults()
	fmt.Printf("%d / %d\n", results.Successes, results.Total)
}
