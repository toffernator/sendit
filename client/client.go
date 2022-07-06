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

func ParseReqs(path string, baseTarget string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		req := parseReq(scanner.Text(), baseTarget)
		addReq(req)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error while parsing requests: %s", err)
	}
}

func parseReq(toParse string, baseTarget string) *models.VerifiableRequest {
	reqOpts := strings.Split(toParse, ",")
	if len(reqOpts) != 4 {
		log.Fatalf("Request '%s' has an invalid format", toParse)
	}

	method := reqOpts[0]
	url := baseTarget + reqOpts[1]
	var body io.Reader
	if reqOpts[2] == "" {
		body = nil
	} else {
		body = strings.NewReader(reqOpts[2])
	}
	var expectedStatusCode int
	if i, err := strconv.Atoi(reqOpts[3]); err == nil {
		expectedStatusCode = i
	} else {
		log.Fatalf("Failed to convert %v to int with err: %s", reqOpts[3], err)
		log.Fatalf("Found a non-integer status code %v", reqOpts[3])
	}

	req := models.NewVerifiableRequest(method, url, body, expectedStatusCode)
	return req
}

func SendReqs() {
	wg := new(sync.WaitGroup)
	for sizeReqs() > 0 {
		wg.Add(1)
		req := removeReq()
		go sendReq(wg, req)
	}
	wg.Wait()
}

func sendReq(wg *sync.WaitGroup, req *models.VerifiableRequest) {
	defer wg.Done()

	resp, err := client.Do(&req.Request)
	if err != nil {
		log.Fatal(err)
	}

	verifiableResponse := models.NewVerifiableResponse(*resp, req.ExpectedStatusCode)
	addResp(verifiableResponse)
}

func ComputeResults() int {
	successes := 0
	for _, resp := range responses {
		if resp.IsSuccessful() {
			successes++
		}
	}
	return successes
}
