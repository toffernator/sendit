package models

import (
	"io"
	"log"
	"net/http"
)

type Job struct {
	Request            *http.Request
	Response           *http.Response
	ExpectedStatusCode int
	Err                error
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
	if j.Response == nil {
		return false
	}

	return j.Response.StatusCode == j.ExpectedStatusCode
}

type Path = string
