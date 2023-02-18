package api

import (
	"github.com/s4kibs4mi/snapify/dao"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/s4kibs4mi/snapify/services"
)

type handlers struct {
	screenshotDao  dao.IScreenshotDao
	tokenDao       dao.ITokenDao
	queueService   services.IQueueService
	storageService services.IBlobStorageService
	logger         log.IAppLogger
}

func NewHandlers(screenshotDao dao.IScreenshotDao, tokenDao dao.ITokenDao,
	queueService services.IQueueService, storageService services.IBlobStorageService, logger log.IAppLogger) IHandlers {
	return &handlers{
		screenshotDao:  screenshotDao,
		tokenDao:       tokenDao,
		queueService:   queueService,
		storageService: storageService,
		logger:         logger,
	}
}
