package tasks

import (
	"bytes"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/jinzhu/gorm"
	"github.com/s4kibs4mi/snapify/app"
	"github.com/s4kibs4mi/snapify/core"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/s4kibs4mi/snapify/models"
	"github.com/s4kibs4mi/snapify/repos"
	"github.com/s4kibs4mi/snapify/services"
	"github.com/s4kibs4mi/snapify/utils"
	"time"
)

const TakeScreenShotTaskName = "take_screen_shot"

func TakeScreenShot(ID string) error {
	repo := repos.NewScreenshotRepo()
	m, err := repo.Get(app.DB(), ID)
	if err != nil {
		log.Log().Errorln(err)

		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return tasks.NewErrRetryTaskLater(err.Error(), time.Second*10)
	}

	result, err := core.TakeScreenShot(m.Website)
	if err != nil {
		log.Log().Errorln(err)
		return tasks.NewErrRetryTaskLater(err.Error(), time.Second*10)
	}

	path := utils.NewUUID() + "-" + utils.FormatUrlWithoutProtocol(m.Website) + ".png"

	if err := services.UploadToMinio(path, "image/png", bytes.NewReader(result), len(result)); err != nil {
		log.Log().Errorln(err)
		return tasks.NewErrRetryTaskLater(err.Error(), time.Second*10)
	}

	m.Status = models.Done
	m.StoredPath = path
	m.UpdatedAt = time.Now()

	if err := repo.Update(app.DB(), m); err != nil {
		log.Log().Errorln(err)
		return tasks.NewErrRetryTaskLater(err.Error(), time.Second*10)
	}
	return nil
}
