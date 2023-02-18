package worker

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/constants"
	"github.com/s4kibs4mi/snapify/dao"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/s4kibs4mi/snapify/services"
)

type IWorker interface {
	Start() error
	Stop() error
	HandleTakeScreenshot(ctx context.Context, task *asynq.Task) error
}

type worker struct {
	server         *asynq.Server
	screenshotDao  dao.IScreenshotDao
	queueService   services.IQueueService
	browserService services.IHeadlessBrowser
	storageService services.IBlobStorageService
	logger         log.IAppLogger
}

func NewWorker(cfg *config.AppCfg, screenshotDao dao.IScreenshotDao, queueService services.IQueueService,
	browserService services.IHeadlessBrowser, storageService services.IBlobStorageService, logger log.IAppLogger) (IWorker, error) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Network: "tcp",
			Addr:    cfg.RedisAddr,
			//Username: cfg.App().RedisUsername,
			Password: cfg.RedisPassword,
			DB:       0,
			//TLSConfig: &tls.Config{
			//	InsecureSkipVerify: true,
			//},
		},
		asynq.Config{
			Concurrency: 1,
			Queues: map[string]int{
				cfg.RedisQueueName: 10,
			},
		},
	)

	return &worker{
		server:         srv,
		screenshotDao:  screenshotDao,
		queueService:   queueService,
		browserService: browserService,
		storageService: storageService,
		logger:         logger,
	}, nil
}

func (w *worker) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(constants.TaskNameHandleTakeScreenshot, w.HandleTakeScreenshot)

	if err := w.server.Start(mux); err != nil {
		return err
	}
	return nil
}

func (w *worker) Stop() error {
	w.server.Stop()
	return nil
}
