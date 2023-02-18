package apimodels

import (
	"strings"
)

type ReqScreenshotCreate struct {
	URL string `json:"url"`
}

func (req *ReqScreenshotCreate) Validate() error {
	req.URL = strings.TrimSpace(req.URL)

	vErr := &ValidationError{}

	if req.URL == "" {
		vErr.Add("url", "is required")
	}

	if vErr.HasErrors() {
		return vErr
	}
	return nil
}
