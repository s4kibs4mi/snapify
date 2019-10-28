package errors

import "encoding/json"

type ValidationError map[string][]string

func (ve *ValidationError) Error() string {
	b, _ := json.Marshal(ve)
	return string(b)
}

func (ve ValidationError) Add(key, value string) {
	ve[key] = append(ve[key], value)
}
