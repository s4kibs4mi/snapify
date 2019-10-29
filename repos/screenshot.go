package repos

import (
	"github.com/jinzhu/gorm"
	"github.com/s4kibs4mi/snapify/models"
	"github.com/s4kibs4mi/snapify/validators"
)

type ScreenshotRepo interface {
	Create(db *gorm.DB, req *validators.ReqCreateScreenshot) ([]models.Screenshot, error)
	Update(db *gorm.DB, m *models.Screenshot) error
	List(db *gorm.DB, page, limit int64) ([]models.Screenshot, error)
	Count(db *gorm.DB) (int, error)
	Search(db *gorm.DB, query string, page, limit int64) ([]models.Screenshot, error)
	Get(db *gorm.DB, ID string) (*models.Screenshot, error)
}
