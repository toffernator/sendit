package generator

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const (
	PATHS = "paths"
)

type PathItem struct {
	Path       string
	Operations []string
}

var api map[string]interface{}

// ParseJSON parse the given JSON file such that other
// methods can run interpreting the API described in that OpenAPI 3.1 file.
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

func Paths() (pathItems []PathItem) {
	paths := api[PATHS].(map[string]interface{})
	pathItems = make([]PathItem, 0)

	for p, ops := range paths {
		operations := ops.(map[string]interface{})

		pathItem := &PathItem{
			Path:       p,
			Operations: make([]string, 0),
		}

		for op := range operations {
			pathItem.Operations = append(pathItem.Operations, op)
		}

		pathItems = append(pathItems, *pathItem)
	}

	return pathItems
}
