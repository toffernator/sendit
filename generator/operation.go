package generator

type Operation struct {
	Verb       string
	Parameters []Parameter
}

type Parameter struct {
	Name     string
	In       string
	Required bool
	Type     string
	Format   string
}

// Operations returns a representation of available verbs and parameters for a given path.
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

func operations(path string) map[string]interface{} {
	paths := paths()
	operations := paths[path].(map[string]interface{})
	return operations
}

func parseParameter(parameters []interface{}) []Parameter {
	if parameters == nil {
		return make([]Parameter, 0)
	}

	parsed := make([]Parameter, 0)

	for _, p := range parameters {
		raw := p.(map[string]interface{})

		Name := raw[PARAMETER_NAME].(string)
		In := raw[PARAMETER_IN].(string)

		Required := false
		if raw[PARAMETER_REQUIRED] != nil {
			Required = raw[PARAMETER_REQUIRED].(bool)
		}

		Type := ""
		if raw[PARAMETER_TYPE] != nil {
			Type = raw[PARAMETER_TYPE].(string)
		}

		Format := ""
		if raw[PARAMETER_FORMAT] != nil {
			Format = raw[PARAMETER_FORMAT].(string)
		}

		parameter := Parameter{
			Name,
			In,
			Required,
			Type,
			Format,
		}

		parsed = append(parsed, parameter)
	}

	return parsed
}
