package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Structs are created using the OpenAPI 3.1 spec:
// https://spec.openapis.org/oas/v3.1.0#info-object

type OpenAPI struct {
	OpenAPI string `json:"openapi"` // The version the OpenApi spec is written in
	Info    Info   `json:"info"`    // Provides metadata about the API
	Paths   map[string]*PathItem
	// TODO: Parse servers to know where to send requests
}

type Info struct {
	Title   string `json:"title"`   // The title of the API
	Summary string `json:"summary"` // A short summary of the API
	Version string `json:"version"` // The version of the OpenAPI document itself
}

type PathItem struct {
	Summary string     `json:"summary"`
	Get     *Operation `json:"get"`
	Put     *Operation `json:"put"`
	Post    *Operation `json:"post"`
	Delete  *Operation `json:"delete"`
	Options *Operation `json:"options"`
	Head    *Operation `json:"head"`
	Trace   *Operation `json:"trace"`
	// TODO: consider parameters?
}

type Operation struct {
	Summary     string `json:"summary"`
	OperationID string `json:"operationId"`
	// TODO: consider parameters
	RequestBody RequestBody `json:"requestBody"`
	// TODO: How to parse the expected response code
	Responses map[string]Response `json:"responses"`
}

type RequestBody struct {
	// Needs special parsing since key may be any of the different content types
	Content  string // Source schema for generating requests from examples here
	Required bool   `json:"required"`
}

type Response struct {
	Description string            `json:"description"`
	Headers     map[string]string // TODO: unstructured data
	Content     map[string]string // TODO: unstructured data
}

func ParseJSON(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Opening json file failed: %v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}
}
