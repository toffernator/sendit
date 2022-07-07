package models

import (
	"io"
	"log"
	"net/http"
)

type Job struct {
	Request            *http.Request
	Response           *http.Response
	HasRun             bool
	ExpectedStatusCode int
}

func NewJob(method string, url string, body io.Reader, expectedStatusCode int) Job {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Error while creating a new request: %s", err)
	}

	return Job{
		Request:            req,
		Response:           nil,
		ExpectedStatusCode: expectedStatusCode,
	}
}

func (j *Job) IsSuccessful() bool {
	if !j.HasRun {
		log.Fatalln("The job must have run before its success can be determined")
	}

	return j.Response.StatusCode == j.ExpectedStatusCode
}
