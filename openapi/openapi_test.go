package generator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toffer/sendit/generator"
)

func TestPathsGivenPetStoreReturns2Pets(t *testing.T) {
	generator.ParseJSON("petstore.json")

	paths := generator.Paths()
	actual := make([]string, 0)
	for _, p := range paths {
		actual = append(actual, p.Path)
	}

	expected := []string{"/pets", "/pets/{petId}"}

	assert.ElementsMatch(t, expected, actual)
}
