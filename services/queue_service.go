package services

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/constants"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/s4kibs4mi/snapify/models"
	"time"
)

type IQueueService interface {
	EnqueueTakeScreenshot(screenshotID uuid.UUID) error
}

type queueService struct {
	queueName   string
	queueClient *asynq.Client
	logger      log.IAppLogger
}

func NewQueueService(cfg *config.AppCfg, logger log.IAppLogger) (IQueueService, error) {
	c := asynq.NewClient(
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
	)

	return &queueService{
		queueName:   cfg.RedisQueueName,
		queueClient: c,
		logger:      logger,
	}, nil
}

func (q *queueService) EnqueueTakeScreenshot(screenshotID uuid.UUID) error {
	q.logger.Info("queuing task")

	params := models.ScreenshotTaskParams{
		ID: screenshotID.String(),
	}
	pld, err := json.Marshal(params)
	if err != nil {
		return err
	}

	task, err := q.queueClient.Enqueue(asynq.NewTask(constants.TaskNameHandleTakeScreenshot, pld,
		asynq.ProcessAt(time.Now().Add(time.Second*5)),
		asynq.Queue(q.queueName)))

	q.logger.Info("Enqueued taskID: ", task.ID)

	return err
}
