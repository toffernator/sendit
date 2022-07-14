package generator

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const (
	PATHS            = "paths"
	PARAMETERS       = "parameters"
	PARAMETER_NAME   = "name"
	PARAMETER_IN     = "in"
	PARAMETER_TYPE   = "type"
	PARAMETER_FORMAT = "format"
)

type PathItem struct {
	Path       string
	Operations []Operation
}

type Operation struct {
	Verb       string
	Parameters []Parameter
}

type Parameter struct {
	Name   string
	In     string
	Type   string
	Format string
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

func Operations(path string) []Operation {
	ops := make([]Operation, 0)

	for v, oitem := range operations(path) {
		params := oitem.(map[string]interface{})[PARAMETERS]
		parameters := make([]Parameter, 0)
		if params != nil {
			parameters = parseParameter(params.([]interface{}))
		}

		operation := Operation{
			Verb:       v,
			Parameters: parameters,
		}

		ops = append(ops, operation)
	}

	return ops
}

func parseParameter(parameters []interface{}) []Parameter {
	if parameters == nil {
		return make([]Parameter, 0)
	}

	parsed := make([]Parameter, 0)
	for _, p := range parameters {
		raw := p.(map[string]interface{})
		parameter := Parameter{
			Name: raw[PARAMETER_NAME].(string),
			In:   raw[PARAMETER_IN].(string),
			Type: raw[PARAMETER_TYPE].(string),
			// TODO coercing potentially nil values
		}

		parsed = append(parsed, parameter)
	}

	return parsed
}

func paths() map[string]interface{} {
	return api[PATHS].(map[string]interface{})
}

func operations(path string) map[string]interface{} {
	paths := paths()
	operations := paths[path].(map[string]interface{})
	return operations
}
