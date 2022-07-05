package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/toffer/sendit/bag"
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
		bag.Add(req)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error while parsing requests: %s", err)
	}
}

func parseReq(req string, baseTarget string) *http.Request {
	reqOpts := strings.Split(req, ",")
	if len(reqOpts) != 4 {
		log.Fatalf("Request '%s' has an invalid format", req)
	}

	method := reqOpts[0]
	url := baseTarget + reqOpts[1]
	var body io.Reader
	if reqOpts[2] == "" {
		body = nil
	} else {
		body = strings.NewReader(reqOpts[2])
	}

	parsedReq, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Failed to create a request with error: %s", err)
	}

	return parsedReq
}

func SendReqs() {
	wg := new(sync.WaitGroup)
	for bag.Size() > 0 {
		wg.Add(1)
		j := bag.Remove()
		go sendReq(wg, j)
	}

	wg.Wait()
}

func sendReq(wg *sync.WaitGroup, req *http.Request) {
	defer wg.Done()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the response here...
	fmt.Println(string(body))
}
