package apimodels

type RespScreenshotData struct {
	ID            string  `json:"id"`
	URL           string  `json:"url"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
	ScreenshotURL *string `json:"screenshot_url,omitempty"`
}

// RespScreenshot represents response body of ScreenshotCreate endpoint
type RespScreenshot struct {
	Data RespScreenshotData `json:"data"`
}

type RespScreenshotList struct {
	Data []RespScreenshotData `json:"data"`
}
