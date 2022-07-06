package models

import (
	"io"
	"log"
	"net/http"
)

type VerifiableRequest struct {
	http.Request
	ExpectedStatusCode int
}

func NewVerifiableRequest(method string, url string, body io.Reader, expectedStatusCode int) *VerifiableRequest {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Error while creating a new request: %s", err)
	}

	return &VerifiableRequest{
		*req,
		expectedStatusCode,
	}
}

type VerifiableResponse struct {
	http.Response
	ExpectedStatusCode int
}

func NewVerifiableResponse(resp http.Response, expectedStatusCode int) *VerifiableResponse {
	return &VerifiableResponse{
		resp,
		expectedStatusCode,
	}
}

func (r *VerifiableResponse) IsSuccessful() bool {
	return r.StatusCode == r.ExpectedStatusCode
}
