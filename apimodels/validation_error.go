package apimodels

import "encoding/json"

type ValidationError struct {
	Errors map[string][]string `json:"errors"`
}

func (vErr *ValidationError) Add(field, msg string) {
	if vErr.Errors == nil {
		vErr.Errors = map[string][]string{}
	}

	vErr.Errors[field] = append(vErr.Errors[field], msg)
}

func (vErr *ValidationError) HasErrors() bool {
	return len(vErr.Errors) > 0
}

func (vErr ValidationError) Error() string {
	b, _ := json.Marshal(vErr)
	return string(b)
}
