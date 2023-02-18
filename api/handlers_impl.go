package api

import (
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/dao"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/s4kibs4mi/snapify/services"
)

type handlers struct {
	cfg            *config.AppCfg
	screenshotDao  dao.IScreenshotDao
	queueService   services.IQueueService
	storageService services.IBlobStorageService
	logger         log.IAppLogger
}

func NewHandlers(cfg *config.AppCfg, screenshotDao dao.IScreenshotDao,
	queueService services.IQueueService, storageService services.IBlobStorageService, logger log.IAppLogger) IHandlers {
	return &handlers{
		cfg:            cfg,
		screenshotDao:  screenshotDao,
		queueService:   queueService,
		storageService: storageService,
		logger:         logger,
	}
}
