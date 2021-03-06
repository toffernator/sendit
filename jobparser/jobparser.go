package jobparser

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/toffer/sendit/models"
)

var (
	Jobs chan models.Job
)

func init() {
	Jobs = make(chan models.Job)
}

// ParseJobs populates a channel with jobs parsed from CSV file given by the path with a base url target,
// closing the channel once done parsing.
func ParseJobs(path string, baseTarget string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		j := parseJob(scanner.Text(), baseTarget)
		Jobs <- j
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error while parsing requests: %s", err)
	}

	close(Jobs)
}

func parseJob(toParse string, baseTarget string) models.Job {
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
