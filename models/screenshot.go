package models

type Status string

const (
	Queued    Status = "queued"
	Failed    Status = "failed"
	Completed Status = "completed"
)

type ScreenshotTaskParams struct {
	ID string `json:"id"`
}
