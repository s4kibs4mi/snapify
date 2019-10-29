package repos

import (
	"github.com/jinzhu/gorm"
	"github.com/s4kibs4mi/snapify/models"
	"github.com/s4kibs4mi/snapify/utils"
	"github.com/s4kibs4mi/snapify/validators"
	"time"
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

func (ssr *ScreenshotRepoImpl) Create(db *gorm.DB, req *validators.ReqCreateScreenshot) ([]models.Screenshot, error) {
	var data []models.Screenshot

	for _, u := range req.URLs {
		m := models.Screenshot{
			ID:        utils.NewUUID(),
			Status:    models.Queued,
			Website:   u,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := db.Model(&m).Create(&m).Error; err != nil {
			return nil, err
		}

		data = append(data, m)
	}
	return data, nil
}

func (ssr *ScreenshotRepoImpl) Update(db *gorm.DB, m *models.Screenshot) error {
	if err := db.Model(m).Save(m).Error; err != nil {
		return err
	}
	return nil
}

func (ssr *ScreenshotRepoImpl) List(db *gorm.DB, page, limit int64) ([]models.Screenshot, error) {
	var data []models.Screenshot
	m := models.Screenshot{}
	if err := db.Model(&m).
		Order("created_at DESC", false).
		Offset((page * limit) - limit).
		Limit(limit).
		Find(&data).Error; err != nil {
		return nil, err
	}

	if len(data) == 0 {
		data = []models.Screenshot{}
	}
	return data, nil
}

func (ssr *ScreenshotRepoImpl) Count(db *gorm.DB) (int, error) {
	var count int
	m := models.Screenshot{}
	if err := db.Model(&m).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (ssr *ScreenshotRepoImpl) Get(db *gorm.DB, ID string) (*models.Screenshot, error) {
	m := models.Screenshot{}
	if err := db.Table(m.TableName()).Where("id = ?", ID).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
