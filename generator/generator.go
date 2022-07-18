package generator

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const (
	PATHS              = "paths"
	PARAMETERS         = "parameters"
	PARAMETER_NAME     = "name"
	PARAMETER_IN       = "in"
	PARAMETER_REQUIRED = "required"
	PARAMETER_TYPE     = "type"
	PARAMETER_FORMAT   = "format"
)

type PathItem struct {
	Path       string
	Operations []Operation
}

var api map[string]interface{}

// ParseJSON parses the given JSON file such that other methods can run interpreting the API described in that OpenAPI
// 3.1 file.
//
// It must be called before any other methods in this package
func ParseJSON(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Opening json file failed: %v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &api)
}

// Paths returns a representation of all available paths and operations to work with them programatically.
func Paths() []PathItem {
	pathItems := make([]PathItem, 0)

	paths := paths()
	for p := range paths {
		operations := Operations(p)
		pathItem := &PathItem{
			Path:       p,
			Operations: operations,
		}

		pathItems = append(pathItems, *pathItem)
	}

	return pathItems
}

func paths() map[string]interface{} {
	return api[PATHS].(map[string]interface{})
}
