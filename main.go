package main

import (
	"fmt"

	"github.com/toffer/sendit/client"
	"github.com/toffer/sendit/jobparser"
)

const (
	REQUESTS = ".local/requests.csv"
	TARGET   = "http://localhost:7777"
)

func main() {
	go jobparser.ParseJobs(REQUESTS, TARGET)
	go client.SendReqs()
	result := client.TallyResults()

	fmt.Printf("%d / %d\n", result.Successes, result.Total)
}
