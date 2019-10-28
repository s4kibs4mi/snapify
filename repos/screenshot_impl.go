package repos

import (
	"github.com/s4kibs4mi/snapify/models"
	"github.com/s4kibs4mi/snapify/validators"
)

type ScreenshotRepoImpl struct {
}

var ssRepo ScreenshotRepo

func NewScreenshotRepo() ScreenshotRepo {
	if ssRepo == nil {
		ssRepo = &ScreenshotRepoImpl{}
	}
	return ssRepo
}

func (ssr *ScreenshotRepoImpl) Create(req *validators.ReqCreateScreenshot) (*models.Screenshot, error) {
	return &models.Screenshot{}, nil
}

func (ssr *ScreenshotRepoImpl) List(page, limit int64) ([]models.Screenshot, error) {
	return nil, nil
}

func (ssr *ScreenshotRepoImpl) Search(query string, page, limit int64) ([]models.Screenshot, error) {
	return nil, nil
}

func (ssr *ScreenshotRepoImpl) Get(ID string) (*models.Screenshot, error) {
	return nil, nil
}
