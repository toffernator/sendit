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
	client.ParseReqs(REQUESTS, TARGET)
	client.SendReqs()
	fmt.Println(client.ComputeResults())
}
