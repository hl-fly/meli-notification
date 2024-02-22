package viewmodel

type ErrorResponse struct {
	Fields []string `json:"errors"`
}

func FormatErrors(errors []error) ErrorResponse {
	errorFields := make([]string, len(errors))
	for i, err := range errors {
		errorFields[i] = err.Error()
	}

	return ErrorResponse{
		Fields: errorFields,
	}
}
