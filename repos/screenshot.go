package repos

import (
	"github.com/s4kibs4mi/snapify/models"
	"github.com/s4kibs4mi/snapify/validators"
)

type ScreenshotRepo interface {
	Create(req *validators.ReqCreateScreenshot) (*models.Screenshot, error)
	List(page, limit int64) ([]models.Screenshot, error)
	Search(query string, page, limit int64) ([]models.Screenshot, error)
	Get(ID string) (*models.Screenshot, error)
}
