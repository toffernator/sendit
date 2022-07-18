package generator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toffer/sendit/generator"
)

func TestOperations(t *testing.T) {
	generator.ParseJSON("petstore.json")
	t.Run("Operations given /pets returns get and post", func(t *testing.T) {
		actual := generator.Operations("/pets")
		expected := []generator.Operation{
			{
				Verb: "get",
				Parameters: []generator.Parameter{
					{
						Name:     "limit",
						In:       "query",
						Required: false,
						Type:     "integer",
						Format:   "int32",
					},
				},
			},
			{
				Verb:       "post",
				Parameters: []generator.Parameter{},
			},
		}

		assert.ElementsMatch(t, expected, actual)
	})
}
