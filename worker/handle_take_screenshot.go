package worker

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/s4kibs4mi/snapify/ent"
	"github.com/s4kibs4mi/snapify/models"
)

func (w *worker) HandleTakeScreenshot(ctx context.Context, task *asynq.Task) error {
	w.logger.Info("Dequeued task: ", task.Type())

	params := models.ScreenshotTaskParams{}
	if err := json.Unmarshal(task.Payload(), &params); err != nil {
		return err
	}

	ss, err := w.screenshotDao.Get(uuid.MustParse(params.ID))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return err
	}

	ssPath, err := w.browserService.TakeScreenshot(ss.URL)
	if err != nil {
		return err
	}

	w.logger.Info("Saved local filepath: ", ssPath)

	storedPath, err := w.storageService.Save(ssPath)
	if err != nil {
		return err
	}

	w.logger.Info("Stored filepath: ", storedPath)

	_, err = ss.Update().
		SetStatus(models.Completed).
		SetStoredPath(storedPath).
		Save(context.Background())
	if err != nil {
		return err
	}

	w.logger.Info("Taking screenshot completed")
	return nil
}
