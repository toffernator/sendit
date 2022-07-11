package main

import (
	"fmt"

	"github.com/toffer/sendit/client"
	"github.com/toffer/sendit/generator"
	"github.com/toffer/sendit/jobparser"
)

const (
	REQUESTS = ".local/requests.csv"
	TARGET   = "http://localhost:7777"
)

func main() {
	go jobparser.ParseJobs(REQUESTS, TARGET)
	go client.SendReqs()
	client.TallyResults()

	for endpoint, result := range client.Results {
		fmt.Printf("%s: %d / %d\n", endpoint, result.Successes, (result.Successes + result.Failures))
	}

	generator.ParseJSON("generator/petstore.json")
}
